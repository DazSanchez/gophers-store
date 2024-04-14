package service

import (
	"log"

	"com.github.dazsanchez/gophers-store/dto"
	"com.github.dazsanchez/gophers-store/repository"
)

type categoryService struct {
	r repository.CategoryRepository
}

var Category categoryService = categoryService{
	r: repository.Category,
}

func (s categoryService) FindAll() dto.CategoriesDTO {
	var cs dto.CategoriesDTO

	m, err := s.r.FindAll()
	if err != nil {
		log.Panicln("can't fetch categories: ", err)
	}

	cs.FromModel(m)

	return cs
}
