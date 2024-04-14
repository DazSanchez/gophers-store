package dto

import "com.github.dazsanchez/gophers-store/model"

// TagDTO is a controller friendly interface for Tag model.
type TagDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// FromModel is a shortcut to populate the reciever dto with data from a given Tag model.
func (dto *TagDTO) FromModel(m model.Tag) {
	*dto = TagDTO{
		Id:   int(m.Id),
		Name: m.Name,
	}
}

// TagDTO is controller friendly interface for a []Tags model.
type TagsDTO []TagDTO

// FromModel is a shortcut to populate the reciever dto with data from a given Tag model slice.
func (dto *TagsDTO) FromModel(ms []model.Tag) {
	ts := make(TagsDTO, 0)
	for _, m := range ms {
		var t TagDTO
		t.FromModel(m)

		ts = append(ts, t)
	}

	*dto = ts
}

// TagRelationDTO is a controller friendly interface for a relation of an Entity with a Tag by its Id.
type TagRelationDTO struct {
	Id int `json:"id"`
}

// ToModel is a shortcut to create a partial Tag model with only the Id field set.
func (dto TagRelationDTO) ToModel() model.Tag {
	return model.Tag{
		Id: int64(dto.Id),
	}
}

// TagsRelationDTO is a controller friendly interface for a collection of relationed Tags to a given Entity by Id.
type TagsRelationDTO []TagRelationDTO

// ToModel is a shortcut to create a slice of Tag model.
func (dtos TagsRelationDTO) ToModel() []model.Tag {
	s := make([]model.Tag, len(dtos))

	for _, dto := range dtos {
		s = append(s, dto.ToModel())
	}

	return s
}

// Ids is a shortcut to retrieve a []int of all tag's Id.
func (dto TagsRelationDTO) Ids() []int {
	ids := make([]int, 0)

	for _, tag := range dto {
		ids = append(ids, tag.Id)
	}

	return ids
}
