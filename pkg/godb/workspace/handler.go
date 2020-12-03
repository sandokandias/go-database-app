package workspace

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sandokandias/go-database-app/pkg/godb"
)

type Handler struct {
	service godb.WorkspaceService
}

func NewHandler(service godb.WorkspaceService) Handler {
	return Handler{service: service}
}

func (h Handler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.create(w, r)
		default:
			http.NotFound(w, r)
		}
	}
}

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	var cw godb.CreateWorkspace
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&cw); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	ctx := r.Context()

	if err := h.service.CreateWorkspace(ctx, cw); err != nil {
		switch e := err.(type) {
		case godb.ValidationError:
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			b, _ := json.Marshal(e)
			w.Write(b)
			return

		default:
			log.Printf("error: %v", e.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
