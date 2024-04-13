package repository

import (
	"database/sql"
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/model"
)

const (
	findAllCategories = "SELECT * FROM category;"
)

type categoryRepository struct{}

var Category categoryRepository = categoryRepository{}

func (r categoryRepository) toCategory(rows *sql.Rows) (model.Category, error) {
	var c model.Category

	err := rows.Scan(&c.Id, &c.Name)

	return c, err
}

func (r categoryRepository) toCategories(rows *sql.Rows) ([]model.Category, error) {
	var cs []model.Category

	for rows.Next() {
		c, err := r.toCategory(rows)
		if err != nil {
			return cs, err
		}

		cs = append(cs, c)
	}

	return cs, nil
}

func (r categoryRepository) FindAll() ([]model.Category, error) {
	rows, err := db.Instance.Query(findAllCategories)

	if err != nil {
		log.Panicln("can't fetch categories: ", err)
		return nil, err
	}

	defer rows.Close()

	return r.toCategories(rows)
}
