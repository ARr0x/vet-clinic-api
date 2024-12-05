package main

import (
	"log"
	"net/http"

	"vet-clinic-api/config"
	"vet-clinic-api/pkg/cat"
	"vet-clinic-api/pkg/treatment"
	"vet-clinic-api/pkg/visit"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	cfg := config.New()
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/cats", cat.Routes(cfg))
		r.Mount("/visits", visit.Routes(cfg))
		r.Mount("/treatments", treatment.Routes(cfg))
	})

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
