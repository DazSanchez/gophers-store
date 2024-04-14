package dto

import "com.github.dazsanchez/gophers-store/model"

// GopherDTO is a controller friendly interface for Gopher model.
type GopherDTO struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Category  CategoryDTO `json:"category"`
	PhotoUrls []string    `json:"photo_urls"`
	Tags      TagsDTO     `json:"tags"`
	Status    string      `json:"status"`
}

// FromModel is a shortcut to populate the reciever dto with data from a given Gopher model.
func (dto *GopherDTO) FromModel(m model.Gopher) {
	var c CategoryDTO
	c.FromModel(m.Category)

	var t TagsDTO
	t.FromModel(m.Tags)

	*dto = GopherDTO{
		Id:        int(m.Id),
		Name:      m.Name,
		Category:  c,
		PhotoUrls: m.PhotoUrls,
		Tags:      t,
		Status:    string(m.Status),
	}
}

// CreateGopherDTO is a controller friendly interface for hanlding data needed to create a new Gopher model.
type CreateGopherDTO struct {
	Name      string               `json:"name"`
	Category  CategoryReferenceDTO `json:"category"`
	PhotoUrls []string             `json:"photoUrls"`
	Tags      TagsRelationDTO      `json:"tags"`
	Status    string               `json:"status"`
}

// ToModel is a shortcut to create an instance of Gopher model.
func (dto CreateGopherDTO) ToModel() model.Gopher {
	return model.Gopher{
		Name:      dto.Name,
		PhotoUrls: dto.PhotoUrls,
		Status:    model.GopherStatus(dto.Status),
		Category:  dto.Category.ToModel(),
		Tags:      dto.Tags.ToModel(),
	}
}
