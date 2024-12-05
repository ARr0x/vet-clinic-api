package treatment

import (
	"encoding/json"
	"net/http"
	"strconv"

	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CreateTreatment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var treatment dbmodel.Treatment
		if err := json.NewDecoder(r.Body).Decode(&treatment); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := treatment.Bind(r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := cfg.TreatmentRepository.Create(&treatment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, treatment)
	}
}

func GetAllTreatments(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		treatments, err := cfg.TreatmentRepository.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, treatments)
	}
}

func GetTreatmentByID(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		treatment, err := cfg.TreatmentRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		render.JSON(w, r, treatment)
	}
}

func UpdateTreatment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var treatment dbmodel.Treatment
		if err := json.NewDecoder(r.Body).Decode(&treatment); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		treatment.ID = uint(id)
		if err := cfg.TreatmentRepository.Update(&treatment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, treatment)
	}
}

func DeleteTreatment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := cfg.TreatmentRepository.Delete(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
