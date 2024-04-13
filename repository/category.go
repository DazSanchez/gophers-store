package repository

import (
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/model"
	"com.github.dazsanchez/gophers-store/query"
	"com.github.dazsanchez/gophers-store/scanner"
)

type categoryRepository struct{}

var Category categoryRepository = categoryRepository{}

func (r categoryRepository) FindAll() ([]model.Category, error) {
	rows, err := db.Instance.Query(query.FindAllCategories)

	if err != nil {
		log.Panicln("can't fetch categories: ", err)
		return nil, err
	}

	defer rows.Close()

	return scanner.ToCategories(rows)
}
