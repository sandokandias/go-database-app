package order

import (
	"reflect"
	"testing"
)

func TestNewOrder(t *testing.T) {
	type test struct {
		id         string
		amount     int64
		items      ItemsData
		customerID string
		want       Order
	}

	tests := []test{
		{
			id:     "order1",
			amount: 3600,
			items: ItemsData{
				ItemData{
					Name:     "Cerveja Budweiser Longneck",
					Price:    300,
					Quantity: 12,
				},
			},
			customerID: "80221189076",
			want: Order{
				ID:     "order1",
				Amount: 3600,
				Items: Items{
					Item{
						ID:       "order1_1",
						Name:     "Cerveja Budweiser Longneck",
						Price:    300,
						Quantity: 12,
					},
				},
				CustomerID: "80221189076",
			},
		},
		{
			id:     "",
			amount: 3600,
			items: ItemsData{
				ItemData{
					Name:     "Cerveja Budweiser Longneck",
					Price:    300,
					Quantity: 12,
				},
			},
			customerID: "80221189076",
			want:       Order{},
		},
		{
			id:     "order2",
			amount: 0,
			items: ItemsData{
				ItemData{
					Name:     "Cerveja Budweiser Longneck",
					Price:    300,
					Quantity: 12,
				},
			},
			customerID: "80221189076",
			want:       Order{},
		},
		{
			id:         "order3",
			amount:     3600,
			items:      ItemsData{},
			customerID: "80221189076",
			want:       Order{},
		},
		{
			id:     "order4",
			amount: 3600,
			items: ItemsData{
				ItemData{
					Name:     "",
					Price:    300,
					Quantity: 12,
				},
			},
			customerID: "80221189076",
			want:       Order{},
		},
		{
			id:     "order5",
			amount: 3600,
			items: ItemsData{
				ItemData{
					Name:     "Cerveja Budweiser Longneck",
					Price:    0,
					Quantity: 12,
				},
			},
			customerID: "80221189076",
			want:       Order{},
		},
		{
			id:     "order6",
			amount: 3600,
			items: ItemsData{
				ItemData{
					Name:     "Cerveja Budweiser Longneck",
					Price:    300,
					Quantity: 0,
				},
			},
			customerID: "80221189076",
			want:       Order{},
		},
	}

	for _, tc := range tests {
		got, _ := NewOrder(tc.id, tc.amount, tc.items, tc.customerID)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
