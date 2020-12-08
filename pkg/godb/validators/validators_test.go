package validators

import (
	"reflect"
	"testing"

	"github.com/sandokandias/go-database-app/pkg/godb"
)

func TestStringRequired(t *testing.T) {
	type test struct {
		field string
		value string
		want  error
	}

	tests := []test{
		{field: "username", value: "admin", want: nil},
		{field: "username", value: "", want: godb.ErrRequiredField("username")},
	}

	for _, tc := range tests {
		got := StringRequired(tc.field, tc.value)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestInt64GreaterZero(t *testing.T) {
	type test struct {
		field string
		value int64
		want  error
	}

	tests := []test{
		{field: "number", value: 42, want: nil},
		{field: "number", value: 0, want: godb.ErrNegativeZeroInt("number")},
		{field: "number", value: -42, want: godb.ErrNegativeZeroInt("number")},
	}

	for _, tc := range tests {
		got := Int64GreaterZero(tc.field, tc.value)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestIntGreaterZero(t *testing.T) {
	type test struct {
		field string
		value int
		want  error
	}

	tests := []test{
		{field: "number", value: 42, want: nil},
		{field: "number", value: 0, want: godb.ErrNegativeZeroInt("number")},
		{field: "number", value: -42, want: godb.ErrNegativeZeroInt("number")},
	}

	for _, tc := range tests {
		got := IntGreaterZero(tc.field, tc.value)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
