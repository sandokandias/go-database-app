package order

import (
	"encoding/json"
	"net/http"

	"github.com/sandokandias/go-database-app/pkg/godb/dhttp/dhttputil"
)

// Handler type that implements http handler
type Handler struct {
	service Service
}

// NewHandler creates a new http handler with service dependency
func NewHandler(service Service) Handler {
	return Handler{service: service}
}

// Handler handles the request and routes the http method to appropriate function
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

// create decodes body to CreateOrder type and invoke Service to run business logic.
// If some error is returned, the handler will render the respose with the detail
func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var co CreateOrder

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&co); err != nil {
		dhttputil.ErrorHandler(err, w, r)
		return
	}

	if err := h.service.CreateOrder(ctx, co); err != nil {
		dhttputil.ErrorHandler(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
