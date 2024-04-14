package service

import (
	"com.github.dazsanchez/gophers-store/repository"
)

type gopherPhotoUrlService struct {
	r repository.GopherPhotoUrlRepository
}

var GopherPhotoUrl gopherPhotoUrlService = gopherPhotoUrlService{
	r: repository.GopherPhotoUrl,
}

func (s gopherPhotoUrlService) FindById(gopherId int64) []string {
	return s.r.FindById(gopherId)
}

func (s gopherPhotoUrlService) AddUrls(gopherId int64, urls []string) []string {
	return s.r.AddUrls(gopherId, urls)
}
