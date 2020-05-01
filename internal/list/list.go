package list

// List is a curated list
type List struct {
	ID          uint
	Name        string
	Description string
	Items       []Item
}

// Item is a list item
type Item struct {
	ID     uint
	Name   string
	URL    string
	PicURL string
}
