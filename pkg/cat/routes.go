package cat

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) *chi.Mux {
	catConfig := New(cfg)
	r := chi.NewRouter()

	r.Post("/", catConfig.CreateCat)
	r.Get("/", GetAllCats(cfg))
	r.Get("/{id}", GetCatByID(cfg))
	r.Put("/{id}", UpdateCat(cfg))
	r.Delete("/{id}", DeleteCat(cfg))

	return r
}
