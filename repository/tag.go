package repository

import (
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/model"
	"com.github.dazsanchez/gophers-store/query"
	"com.github.dazsanchez/gophers-store/scanner"
)

type tagRepository struct{}

var Tag tagRepository = tagRepository{}

func (r tagRepository) FindAll() ([]model.Tag, error) {
	rows, err := db.Instance.Query(query.FindAllTags)

	if err != nil {
		log.Panicln("can't fetch tags: ", err)
		return nil, err
	}

	defer rows.Close()

	return scanner.ToTags(rows)
}
