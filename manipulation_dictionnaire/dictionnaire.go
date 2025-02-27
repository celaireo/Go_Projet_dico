package manipulation_dictionnaire

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"sync"
)

// Structure d'un mot du dictionnaire
type DictionaryEntry struct {
	Mot        string `json:"mot"`
	Definition string `json:"definition"`
}

// Structure du dictionnaire avec une protection contre les accès concurrents
type Dictionary struct {
	entries map[string]string
	mu      sync.RWMutex
}

// Initialisation du dictionnaire
func NewDictionary() *Dictionary {
	return &Dictionary{
		entries: make(map[string]string),
	}
}

// Ajouter un mot au dictionnaire
func (d *Dictionary) Add(w http.ResponseWriter, r *http.Request) {
	var entry DictionaryEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}

	// ✅ Vérification : Le mot et la définition ne doivent pas être vides
	if strings.TrimSpace(entry.Mot) == "" || strings.TrimSpace(entry.Definition) == "" {
		http.Error(w, "Le mot et la définition sont obligatoires", http.StatusBadRequest)
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()
	d.entries[entry.Mot] = entry.Definition

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Mot ajouté avec succès"})
}

// Modifier un mot existant
func (d *Dictionary) Update(w http.ResponseWriter, r *http.Request) {
	var entry DictionaryEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.entries[entry.Mot]; !exists {
		http.Error(w, "Mot non trouvé", http.StatusNotFound)
		return
	}

	d.entries[entry.Mot] = entry.Definition
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Mot modifié avec succès"})
}

// Supprimer un mot du dictionnaire
func (d *Dictionary) Remove(w http.ResponseWriter, r *http.Request) {
	mot := r.URL.Query().Get("mot")
	if mot == "" {
		http.Error(w, "Paramètre 'mot' manquant", http.StatusBadRequest)
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.entries[mot]; !exists {
		http.Error(w, "Mot non trouvé", http.StatusNotFound)
		return
	}

	delete(d.entries, mot)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Mot supprimé avec succès"})
}

// Supprimer tous les mots du dictionnaire
func (d *Dictionary) RemoveAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	// ✅ Réinitialise complètement le dictionnaire
	d.entries = make(map[string]string)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tous les mots ont été supprimés"})
}

// Liste de tous les mots
func (d *Dictionary) List(w http.ResponseWriter, r *http.Request) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	var results []DictionaryEntry
	for mot, definition := range d.entries {
		results = append(results, DictionaryEntry{Mot: mot, Definition: definition})
	}

	json.NewEncoder(w).Encode(results)
}

// Recherche avancée d'un mot
func (d *Dictionary) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Paramètre 'query' manquant", http.StatusBadRequest)
		return
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	var results []DictionaryEntry
	for mot, definition := range d.entries {
		if strings.Contains(strings.ToLower(mot), strings.ToLower(query)) {
			results = append(results, DictionaryEntry{Mot: mot, Definition: definition})
		}
	}

	if len(results) == 0 {
		http.Error(w, "Aucun mot trouvé", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(results)
}

// Nombre total de mots dans le dictionnaire
func (d *Dictionary) Count(w http.ResponseWriter, r *http.Request) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	json.NewEncoder(w).Encode(map[string]int{"total_mots": len(d.entries)})
}

// Vérification de l'état du serveur
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Sauvegarder le dictionnaire dans un fichier JSON
func (d *Dictionary) SaveToFile(filename string) error {
	d.mu.RLock()
	defer d.mu.RUnlock()

	data, err := json.MarshalIndent(d.entries, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Charger les mots à partir d'un fichier JSON
func (d *Dictionary) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	return json.Unmarshal(data, &d.entries)
}
