package repository

import (
	"database/sql"
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/model"
)

const (
	findAllTags = "SELECT * FROM tag;"
)

type tagRepository struct{}

var Tag tagRepository = tagRepository{}

func (r tagRepository) toTag(rows *sql.Rows) (model.Tag, error) {
	var t model.Tag

	err := rows.Scan(&t.Id, &t.Name)

	return t, err
}

func (r tagRepository) toTags(rows *sql.Rows) ([]model.Tag, error) {
	var ts []model.Tag

	for rows.Next() {
		c, err := r.toTag(rows)

		if err != nil {
			return ts, err
		}

		ts = append(ts, c)
	}

	return ts, nil
}

func (r tagRepository) FindAll() ([]model.Tag, error) {
	rows, err := db.Instance.Query(findAllTags)

	if err != nil {
		log.Panicln("can't fetch tags: ", err)
		return nil, err
	}

	defer rows.Close()

	return r.toTags(rows)
}
