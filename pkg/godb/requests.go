package godb

// CreateOrder type that represents a request for order creation
type CreateOrder struct {
	ID     string `json:"id"`
	Amount int64  `json:"amount"`
}
