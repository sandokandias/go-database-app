package validators

import (
	"github.com/sandokandias/go-database-app/pkg/godb"
)

// RequiredString validates the required string
func RequiredString(field, value string) error {
	if value == "" {
		return godb.ErrRequiredField(field)
	}
	return nil
}
