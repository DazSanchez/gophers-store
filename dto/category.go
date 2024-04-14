package dto

import "com.github.dazsanchez/gophers-store/model"

// CategoryDTO is a controller friendly interface for Category model.
type CategoryDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// FromModel is a shortcut to populate the reciever dto with data from a given Category model.
func (dto *CategoryDTO) FromModel(m model.Category) {
	*dto = CategoryDTO{
		Id:   int(m.Id),
		Name: m.Name,
	}
}

// CategoriesDTO is controller friendly interface for a []Category model.
type CategoriesDTO []CategoryDTO

// FromModel is a shortcut to populate the reciever dto with data from a given []Category model.
func (dto *CategoriesDTO) FromModel(ms []model.Category) {
	cs := make(CategoriesDTO, 0)

	for _, m := range ms {
		var t CategoryDTO
		t.FromModel(m)

		cs = append(cs, t)
	}

	*dto = cs
}

// CategoryReferenceDTO is a controller friendly interface that represents a relation of an Entity with a Category by its Id.
type CategoryReferenceDTO struct {
	Id int `json:"id"`
}

// ToModel is a shortcut to create a partial Category model with only the Id field set.
func (dto CategoryReferenceDTO) ToModel() model.Category {
	return model.Category{
		Id: int64(dto.Id),
	}
}
