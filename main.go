package main

import (
	"Go_Project_Dico/manipulation_dictionnaire"
	"Go_Project_Dico/server" // ✅ Import du package server
	"log"
	"net/http"
	"time"
)

const port = 8080

func main() {
	dictionary := manipulation_dictionnaire.NewDictionary()

	// Charger les données sauvegardées
	if err := dictionary.LoadFromFile("data/dico.json"); err != nil {
		log.Println("⚠️ Aucune donnée trouvée, dictionnaire vide.")
	}

	mux := http.NewServeMux()
	manipulation_dictionnaire.SetupRoutes(mux, dictionary)

	// Ajouter les middlewares
	muxWithLogging := loggingMiddleware(mux)
	muxWithCORS := corsMiddleware(muxWithLogging)

	serverInstance := &http.Server{
		Addr:         ":8080",
		Handler:      muxWithCORS,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("🚀 Serveur démarré sur http://localhost:%d\n", port)

	// ✅ Utilisation correcte de gracefulShutdown depuis server.go
	go server.GracefulShutdown(serverInstance, dictionary)

	if err := serverInstance.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("❌ Erreur serveur: %s", err)
	}
}

// Middleware pour journaliser les requêtes
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Middleware CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
