package server

import (
	"Go_Project_Dico/manipulation_dictionnaire"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ✅ Fonction exportée (nom avec majuscule)
func GracefulShutdown(server *http.Server, dictionary *manipulation_dictionnaire.Dictionary) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("🛑 Arrêt du serveur en cours...")

	// Sauvegarde des données avant arrêt
	if err := dictionary.SaveToFile("data/dico.json"); err != nil {
		log.Println("❌ Erreur lors de la sauvegarde:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Erreur lors de l'arrêt: %s", err)
	}

	log.Println("✅ Serveur arrêté proprement.")
}
