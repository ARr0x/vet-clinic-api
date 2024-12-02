package visit

import (
	"net/http"
	"strconv"

	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CreateVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var visit dbmodel.Visit
		if err := render.Bind(r, &visit); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := cfg.VisitRepository.Create(&visit); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visit)
	}
}

func GetVisitByID(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		visit, err := cfg.VisitRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, "Visit not found", http.StatusNotFound)
			return
		}
		render.JSON(w, r, visit)
	}
}

func GetVisitsByCatID(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		catID, err := strconv.Atoi(chi.URLParam(r, "catID"))
		if err != nil {
			http.Error(w, "Invalid Cat ID", http.StatusBadRequest)
			return
		}

		visits, err := cfg.VisitRepository.FindByCatID(uint(catID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visits)
	}
}

func UpdateVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var visit dbmodel.Visit
		if err := render.Bind(r, &visit); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		visit.ID = uint(id)

		if err := cfg.VisitRepository.Update(&visit); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visit)
	}
}

func DeleteVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := cfg.VisitRepository.Delete(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
