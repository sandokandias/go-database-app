package godb

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ValidationError type that represents a validation error
type ValidationError struct {
	Field string
	Err   error
}

// ErrRequiredField error that represents a required field
func ErrRequiredField(field string) ValidationError {
	return ValidationError{Field: field,
		Err: errors.New("field is required")}
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation: %s", v.Err.Error())
}

func (v ValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Field string `json:"field"`
		Error string `json:"error"`
	}{
		Field: v.Field,
		Error: v.Error(),
	})
}
