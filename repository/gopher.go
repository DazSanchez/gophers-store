package repository

import (
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/model"
	"com.github.dazsanchez/gophers-store/query"
	"com.github.dazsanchez/gophers-store/scanner"
)

type GopherRepository struct{}

// Gopher allows to manipulate database's gopher table.
var Gopher GopherRepository = GopherRepository{}

// FindById retrieves a record that matches the given gopherId as []Gopher model.
// It panics if can't retrieve the data or can't parse to []Gopher.
func (r GopherRepository) FindById(id int64) model.Gopher {
	gRow := query.FindGopherById(id).RunWith(db.Instance).QueryRow()
	g, err := scanner.ToGopher(gRow)

	if err != nil {
		log.Panicln("can't parse gopher: ", err)
	}

	return g
}

// Create creates a record based on the given Gopher model.
// It panics if can't retrieve insert data or can't parse the result into Gopher.
func (r GopherRepository) Create(g model.Gopher) model.Gopher {
	defer func() {
		if r := recover(); r != nil {
			log.Panicln("can't insert gopher: ", r)
		}
	}()

	rs, err := query.CreateGopher(g).RunWith(db.Instance).Exec()
	if err != nil {
		log.Panicln(err)
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Panicln("can't get last id")
	}

	return r.FindById(id)
}
