package service

import (
	"log"

	"com.github.dazsanchez/gophers-store/db"
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

	tx, err := db.TxBegin()
	if err != nil {
		log.Panicln("can't start transaction: ", err)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("recover from panic, rolling back Tx: ", r)
			tx.Rollback()
		}
	}()

	m := s.gr.CreateWithTx(tx, src.ToModel())
	m.PhotoUrls = s.gpu.AddUrlsWithTx(tx, m.Id, src.PhotoUrls)
	m.Tags = s.gt.AddTagsWithTx(tx, m.Id, src.Tags.Ids())

	err = tx.Commit()
	if err != nil {
		log.Panicln("can't commit transaction: ", err)
	}

	g.FromModel(m)

	return g
}
