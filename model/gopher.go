package model

// GopherStatus represents the availability of a given Gopher in the store.
type GopherStatus string

const (
	StatusAvailable GopherStatus = "available"
	StatusPending   GopherStatus = "pending"
	StatusSold      GopherStatus = "sold"
)

// Gopher represents an entity that can be sell in the store.
// PhotoUrls and Tags could be empty.
type Gopher struct {
	Id        int64
	Name      string
	Category  Category
	PhotoUrls []string
	Tags      []Tag
	Status    GopherStatus
}
