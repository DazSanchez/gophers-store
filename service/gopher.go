package service

import (
	"com.github.dazsanchez/gophers-store/dto"
	"com.github.dazsanchez/gophers-store/repository"
)

type gopherService struct {
	gpu repository.GopherPhotoUrlRepository
	gr  repository.GopherRepository
	gt  repository.GopherTagRepository
}

var Gopher gopherService = gopherService{
	gpu: repository.GopherPhotoUrl,
	gr:  repository.Gopher,
	gt:  repository.GopherTag,
}

func (s gopherService) FindById(id int64) dto.GopherDTO {
	var g dto.GopherDTO

	m := s.gr.FindById(id)
	m.PhotoUrls = s.gpu.FindById(m.Id)
	m.Tags = s.gt.FindById(m.Id)

	g.FromModel(m)

	return g
}

func (s gopherService) Create(src dto.CreateGopherDTO) dto.GopherDTO {
	var g dto.GopherDTO

	m := s.gr.Create(src.ToModel())
	m.PhotoUrls = s.gpu.AddUrls(m.Id, src.PhotoUrls)
	m.Tags = s.gt.AddTags(m.Id, src.Tags.Ids())

	g.FromModel(m)

	return g
}
