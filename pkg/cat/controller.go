package cat

import (
	"encoding/json"
	"net/http"
	"strconv"

	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CreateCat(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cat dbmodel.Cat
		if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := cat.Bind(r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := cfg.CatRepository.Create(&cat); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, cat)
	}
}

func GetAllCats(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cats, err := cfg.CatRepository.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, cats)
	}
}

func GetCatByID(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		cat, err := cfg.CatRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		render.JSON(w, r, cat)
	}
}

func UpdateCat(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var cat dbmodel.Cat
		if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cat.ID = uint(id)
		if err := cfg.CatRepository.Update(&cat); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, cat)
	}
}

func DeleteCat(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := cfg.CatRepository.Delete(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
