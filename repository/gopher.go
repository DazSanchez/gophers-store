package repository

import (
	"database/sql"
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/model"
	"com.github.dazsanchez/gophers-store/query"
	"com.github.dazsanchez/gophers-store/scanner"
	sq "github.com/Masterminds/squirrel"
)

type GopherRepository struct{}

// Gopher allows to manipulate database's gopher table.
var Gopher GopherRepository = GopherRepository{}

func (r GopherRepository) findById(id int64, runner sq.BaseRunner) model.Gopher {
	gRow := query.FindGopherById(id).RunWith(runner).QueryRow()
	g, err := scanner.ToGopher(gRow)

	if err != nil {
		log.Panicln("can't parse gopher: ", err)
	}

	return g
}

// FindById retrieves a record that matches the given gopherId as []Gopher model.
// It panics if can't retrieve the data or can't parse to []Gopher.
func (r GopherRepository) FindById(id int64) model.Gopher {
	return r.findById(id, db.Instance)
}

func (r GopherRepository) create(g model.Gopher, runner sq.BaseRunner) model.Gopher {
	defer func() {
		if r := recover(); r != nil {
			log.Panicln("can't insert gopher: ", r)
		}
	}()

	rs, err := query.CreateGopher(g).RunWith(runner).Exec()
	if err != nil {
		log.Panicln(err)
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Panicln("can't get last id")
	}

	return r.findById(id, runner)
}

// Create creates a record based on the given Gopher model.
// It panics if can't retrieve insert data or can't parse the result into Gopher.
func (r GopherRepository) Create(g model.Gopher) model.Gopher {
	return r.create(g, db.Instance)
}

// CreateWithTx creates a record based on the given Gopher model.
// Allows to run the query inside a Transaction context.
// It panics if can't retrieve insert data or can't parse the result into Gopher.
func (r GopherRepository) CreateWithTx(tx *sql.Tx, g model.Gopher) model.Gopher {
	return r.create(g, tx)
}
