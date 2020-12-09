package order

import (
	"context"
	"reflect"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	type test struct {
		ctx         context.Context
		createOrder CreateOrder
		want        error
	}

	tests := []test{
		{
			ctx: context.Background(),
			createOrder: CreateOrder{
				ID:     "order1",
				Amount: 3600,
				Items: ItemsData{
					ItemData{
						Name:     "Cerveja Budweiser Longneck",
						Price:    300,
						Quantity: 12,
					},
				},
				Customer: CustomerData{
					Name:     "Mary Doo",
					Document: "80221189076",
					Address:  "Rua Uber, 100",
				},
			},
			want: nil,
		},
	}

	for _, tc := range tests {
		service := NewService(
			MockTxManager{},
			MockOrderStorage{},
			MockCustomerStorage{},
		)

		got := service.CreateOrder(tc.ctx, tc.createOrder)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
