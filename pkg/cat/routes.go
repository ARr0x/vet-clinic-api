package cat

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Post("/", CreateCat(cfg))
	r.Get("/", GetAllCats(cfg))
	r.Get("/{id}", GetCatByID(cfg))
	r.Put("/{id}", UpdateCat(cfg))
	r.Delete("/{id}", DeleteCat(cfg))
	return r
}
