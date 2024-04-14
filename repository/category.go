package repository

import (
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/model"
	"com.github.dazsanchez/gophers-store/query"
	"com.github.dazsanchez/gophers-store/scanner"
)

type CategoryRepository struct{}

// Category allows to manipulate database's category table.
var Category CategoryRepository = CategoryRepository{}

// FindAll retrieves all records from the repository as []Category model.
// It panics if can't retrieve data or can't parse into []Category model.
func (r CategoryRepository) FindAll() ([]model.Category, error) {
	rows, err := query.FindAllCategories.RunWith(db.Instance).Query()

	if err != nil {
		log.Panicln("can't fetch categories: ", err)
	}

	defer rows.Close()

	return scanner.ToCategories(rows)
}
