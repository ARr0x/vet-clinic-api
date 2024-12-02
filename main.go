package main

import (
	"log"
	"net/http"
	"os"

	"vet-clinic-api/config"
	"vet-clinic-api/pkg/cat"
	"vet-clinic-api/pkg/treatment"
	"vet-clinic-api/pkg/visit"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	// Charger les variables d'environnement
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN non défini")
	}

	// Initialiser la configuration
	cfg := config.New(dsn)

	// Configurer le routeur
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Routes principales
	r.Mount("/api/v1/cats", cat.Routes(cfg))
	r.Mount("/api/v1/visits", visit.Routes(cfg))
	r.Mount("/api/v1/treatments", treatment.Routes(cfg))

	// Démarrer le serveur
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Serveur démarré sur le port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
