package service

import (
	"log"

	"com.github.dazsanchez/gophers-store/dto"
	"com.github.dazsanchez/gophers-store/repository"
)

type tagService struct {
	r repository.TagRepository
}

var Tag tagService = tagService{
	r: repository.Tag,
}

func (s tagService) FindAll() dto.TagsDTO {
	var ts dto.TagsDTO

	m, err := s.r.FindAll()
	if err != nil {
		log.Panicln("can't fetch tags: ", err)
	}

	ts.FromModel(m)

	return ts
}
