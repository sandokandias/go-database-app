package handlerutil

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sandokandias/go-database-app/pkg/godb"
)

// ErrorHandler handles error in handler layer and encodes to appropriate type
func ErrorHandler(err error, w http.ResponseWriter, r *http.Request) {
	switch e := err.(type) {
	case godb.ValidationError:
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(e)
		w.Write(b)
	default:
		log.Printf("error request handler: %v", e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
