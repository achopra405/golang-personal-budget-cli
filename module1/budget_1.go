package module1

// Budget stores budget information
type Budget struct {
	Items []Item
}

// Item stores item information
type Item struct {
	Description string
	Value       float32
}
