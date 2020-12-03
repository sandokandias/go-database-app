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

// Int64GreaterZero validates the int value that needs to be greater than zero
func Int64GreaterZero(field string, value int64) error {
	if value <= 0 {
		return godb.ErrNegativeZeroInt(field)
	}
	return nil
}

// IntGreaterZero validates the int value that needs to be greater than zero
func IntGreaterZero(field string, value int) error {
	if value <= 0 {
		return godb.ErrNegativeZeroInt(field)
	}
	return nil
}
