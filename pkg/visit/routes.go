package visit

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	r.Post("/", CreateVisit(cfg))
	r.Get("/{id}", GetVisitByID(cfg))
	r.Get("/cat/{catID}", GetVisitsByCatID(cfg))
	r.Put("/{id}", UpdateVisit(cfg))
	r.Delete("/{id}", DeleteVisit(cfg))

	return r
}
