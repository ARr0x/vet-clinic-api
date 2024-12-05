package visit

import (
	"encoding/json"
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
		if err := json.NewDecoder(r.Body).Decode(&visit); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := visit.Bind(r); err != nil {
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

func GetAllVisits(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		visits, err := cfg.VisitRepository.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visits)
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
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		render.JSON(w, r, visit)
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
		if err := json.NewDecoder(r.Body).Decode(&visit); err != nil {
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
