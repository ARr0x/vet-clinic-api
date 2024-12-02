package cat

import (
	"net/http"
	"strconv"

	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CatConfig struct {
    *config.Config
}

func New(configuration *config.Config) *CatConfig {
    return &CatConfig{configuration}
}

func (c *CatConfig) CreateCat(w http.ResponseWriter, r *http.Request) {
    // Étape 1 : Lire et valider la requête
    var req models.CatRequest
    if err := render.Bind(r, &req); err != nil {
        render.JSON(w, r, map[string]string{"error": "Invalid request"})
        return
    }

    // Étape 2 : Préparer l'objet à enregistrer
    cat := &dbmodel.Cat{
        Name:  req.Name,
        Age:   req.Age,
        Breed: req.Breed,
    }

    // Étape 3 : Créer le chat dans la base de données
    err := c.CatRepository.Create(cat)
    if err != nil {
        render.JSON(w, r, map[string]string{"error": "Failed to create cat"})
        return
    }

    // Étape 4 : Construire la réponse
    resp := models.CatResponse{
        ID:     cat.ID,
        Name:   cat.Name,
        Age:    cat.Age,
        Breed:  cat.Breed,
        Weight: cat.Weight,
    }

    // Étape 5 : Envoyer la réponse
    render.JSON(w, r, resp)
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

        var req models.CatRequest
        if err := render.Bind(r, &req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Vous pouvez maintenant utiliser les données de req pour mettre à jour le chat
        cat := dbmodel.Cat{
            ID:     uint(id),
            Name:   req.Name,
            Age:    req.Age,
            Breed:  req.Breed,
        }

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
