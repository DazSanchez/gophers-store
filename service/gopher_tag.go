package service

import (
	"com.github.dazsanchez/gophers-store/dto"
	"com.github.dazsanchez/gophers-store/repository"
)

type gopherTagService struct {
	r repository.GopherTagRepository
}

var GopherTag gopherTagService = gopherTagService{
	r: repository.GopherTag,
}

func (s gopherTagService) FindById(gopherId int64) dto.TagsDTO {
	var ts dto.TagsDTO

	m := s.r.FindById(gopherId)
	ts.FromModel(m)

	return ts
}

func (s gopherTagService) AddTags(gopherId int64, tags dto.TagsRelationDTO) dto.TagsDTO {
	var ts dto.TagsDTO

	m := s.r.AddTags(gopherId, tags.Ids())

	ts.FromModel(m)

	return ts
}
