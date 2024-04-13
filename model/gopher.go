package model

type GopherStatus string

const (
	StatusAvailable GopherStatus = "available"
	StatusPending   GopherStatus = "pending"
	StatusSold      GopherStatus = "sold"
)

type Gopher struct {
	Id        int64        `json:"id"`
	Name      string       `json:"name"`
	Category  Category     `json:"category"`
	PhotoUrls []string     `json:"photoUrls"`
	Tags      []Tag        `json:"tags"`
	Status    GopherStatus `json:"status"`
}
