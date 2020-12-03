package godb

// CreateOrder type that represents a request for order creation
type CreateOrder struct {
	ID     string     `json:"id"`
	Amount int64      `json:"amount"`
	Items  []ItemData `json:"items"`
}

// ItemData type that represents a item of the order entity
type ItemData struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}
