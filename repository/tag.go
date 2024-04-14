package repository

import (
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/model"
	"com.github.dazsanchez/gophers-store/query"
	"com.github.dazsanchez/gophers-store/scanner"
)

type TagRepository struct{}

// Tag allows to manipulate database's tag table.
var Tag TagRepository = TagRepository{}

// FindAll retrieves all records from the repository as []Tag model.
// It panics if can't retrieve data or can't parse into []Tag model.
func (r TagRepository) FindAll() ([]model.Tag, error) {
	rows, err := query.FindAllTags.RunWith(db.Instance).Query()

	if err != nil {
		log.Panicln("can't fetch tags: ", err)
	}

	defer rows.Close()

	return scanner.ToTags(rows)
}
