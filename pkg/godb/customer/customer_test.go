package customer

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type test struct {
		name     string
		document string
		address  string
		want     Customer
	}

	tests := []test{
		{
			name:     "Mary",
			document: "80221189076",
			address:  "Rua Uber, 100",
			want: Customer{
				Name:     "Mary",
				Document: "80221189076",
				Address:  "Rua Uber, 100",
			},
		},
		{
			name:     "",
			document: "80221189076",
			address:  "Rua Uber, 100",
			want:     Customer{},
		},
		{
			name:     "Mary",
			document: "",
			address:  "Rua Uber, 100",
			want:     Customer{},
		},
		{
			name:     "Mary",
			document: "80221189076",
			address:  "",
			want:     Customer{},
		},
		{
			name:     "",
			document: "",
			address:  "",
			want:     Customer{},
		},
	}

	for _, tc := range tests {
		got, _ := New(tc.name, tc.document, tc.address)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
