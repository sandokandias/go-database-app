package order

import (
	"encoding/json"
	"net/http"

	"github.com/sandokandias/go-database-app/pkg/godb/handler/handlerutil"

	"github.com/sandokandias/go-database-app/pkg/godb"
)

// Handler type that implements http handler
type Handler struct {
	service godb.OrderService
}

// NewHandler creates a new http handler with service dependency
func NewHandler(service godb.OrderService) Handler {
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

// create decode body to CreateOrder type and invoke OrderService to run business rules.
// If some error is returned, the handler will render the body with the detail
func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var co godb.CreateOrder

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&co); err != nil {
		handlerutil.ErrorHandler(err, w, r)
		return
	}

	if err := h.service.CreateOrder(ctx, co); err != nil {
		handlerutil.ErrorHandler(err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
