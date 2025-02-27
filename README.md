# 📖 **Projet Dictionnaire en Go**
🚀 **Dictionnaire en Go 1.24.0** : Une API REST permettant d'ajouter, modifier, rechercher et supprimer des mots dans un dictionnaire, avec persistance des données et gestion des accès concurrents.

📌 **Développé par** : **Groupe 4**  
- OKA Celaire
- Salamata Nourou MBAYE
- Rostom MOUADDEB


---

## 🌟 **Fonctionnalités**
✅ **Ajout de mots avec leur définition**  
✅ **Modification des définitions**  
✅ **Suppression de mots**  
✅ **Liste complète des mots enregistrés**  
✅ **Recherche avancée de mots (supporte la recherche partielle)**  
✅ **Compteur du nombre total de mots**  
✅ **Persistance des données (fichier JSON)**  
✅ **Gestion avancée des routes et méthodes HTTP**  
✅ **Arrêt propre du serveur avec sauvegarde automatique**  
✅ **Middleware CORS et Logging**  
✅ **Route `/health` pour vérifier l’état du serveur**

---

## 🔧 **Prérequis**
Avant de commencer, assure-toi d’avoir installé **Go 1.24.0** ou une version supérieure.

**Vérifier la version installée :**
```sh
go version
```

---

## 🚀 **Installation et Démarrage**
### **1️⃣ Cloner le dépôt**
```sh
git clone https://github.com/celaireo/Go_Project_Dico.git
cd Go_Project_Dico
```

### **2️⃣ Initialiser les dépendances**
```sh
go mod tidy
```

### **3️⃣ Lancer le serveur**
```sh
go run main.go
```
Le serveur est accessible à l’adresse **http://localhost:8080**.

---

## 📡 **Routes API**
### 📍 **Routes CRUD**
| Méthode | Route        | Description |
|---------|-------------|-------------|
| **POST** | `/add`       | Ajoute un mot avec sa définition |
| **PUT**  | `/update`    | Modifie un mot existant |
| **DELETE** | `/remove?mot=<mot>` | Supprime un mot |
| **GET**  | `/list`      | Liste tous les mots |

---

### 📍 **Routes avancées**
| Méthode | Route        | Description |
|---------|-------------|-------------|
| **GET**  | `/search?query=<mot>` | Recherche un mot (supporte la recherche partielle) |
| **GET**  | `/count`      | Compte le nombre total de mots |
| **GET**  | `/health`     | Vérifie l’état du serveur |

---

## 🛠️ **Exemples d’Utilisation**
### **Ajouter un mot**
```sh
curl -X POST http://localhost:8080/add -d '{"mot":"chat", "definition":"Petit félin domestique"}' -H "Content-Type: application/json"
```

### **Modifier un mot**
```sh
curl -X PUT http://localhost:8080/update -d '{"mot":"chat", "definition":"Félin domestique affectueux"}' -H "Content-Type: application/json"
```

### **Rechercher un mot (supporte la recherche partielle)**
```sh
curl -X GET "http://localhost:8080/search?query=cha"
```

### **Supprimer un mot**
```sh
curl -X DELETE "http://localhost:8080/remove?mot=chat"
```

### **Lister tous les mots**
```sh
curl -X GET "http://localhost:8080/list"
```

---

## 💾 **Persistance des Données**
Les mots sont enregistrés dans le fichier **`data/dico.json`**.  
Toutes les modifications sont automatiquement sauvegardées.

**💡 Reprise après arrêt :**  
Si le serveur est redémarré, il recharge les données sauvegardées.

---

## 🛑 **Arrêt propre du serveur**
Lorsque le serveur est arrêté (`CTRL+C`), une **sauvegarde automatique** est effectuée.  
Toutes les données sont enregistrées dans **`data/dico.json`** avant l'arrêt.

---

## 💡 **Comment Contribuer ?**
Les contributions sont **les bienvenues** ! 🎉

### 📌 **1. Forker le projet**
```sh
git clone https://github.com/
```

### 📌 **2. Créer une nouvelle branche**
```sh
git checkout -b feature-nouvelle-fonctionnalité
```

### 📌 **3. Ajouter tes modifications**
```sh
git add .
git commit -m "Ajout de la fonctionnalité XYZ"
```

### 📌 **4. Pousser les modifications**
```sh
git push origin feature-nouvelle-fonctionnalité
```

### 📌 **5. Créer une Pull Request**
Va sur GitHub et propose ta modification via une **Pull Request**.

---

## 🛠 **Technologies Utilisées**
- **Langage :** Go 1.24.0
- **Base de données :** Fichier JSON (persistance des mots)
- **Framework HTTP :** `net/http`
- **Gestion des concurrents :** `sync.RWMutex`

---

## 📝 **Licence**
Ce projet est sous licence **MIT**.  
Tu es libre de l’utiliser, le modifier et le distribuer. 😊

---

## 👨‍💻 **Développeur**

🔥 **Merci d’utiliser le projet Dictionnaire en Go !** 🚀  

