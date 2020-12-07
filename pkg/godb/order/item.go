package order

// Items type that represents a collection of items
type Items []Item

// Item type that represents a item of the order entity
type Item struct {
	ID       string
	Name     string
	Price    int64
	Quantity int
}
