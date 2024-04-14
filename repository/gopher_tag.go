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

type GopherTagRepository struct{}

// GopherTag allows to manipulate database's gopher_tag table.
var GopherTag GopherTagRepository = GopherTagRepository{}

func (r GopherTagRepository) findById(gopherId int64, runner sq.BaseRunner) []model.Tag {
	tRows, err := query.FindGopherTags(gopherId).RunWith(runner).Query()
	if err != nil {
		log.Panicln("can't fetch gopher tags: ", err)
	}

	defer tRows.Close()

	tags, err := scanner.ToTags(tRows)
	if err != nil {
		log.Panicln("can't parse gopher tags")
	}

	return tags
}

// FindById retrieves all records that matches the given gopherId as []Tag model.
// It panics if can't retrieve the data or can't parse to []Tag.
func (r GopherTagRepository) FindById(gopherId int64) []model.Tag {
	return r.findById(gopherId, db.Instance)
}

func (r GopherTagRepository) addTags(gopherId int64, tagIds []int, runner sq.BaseRunner) []model.Tag {
	_, err := query.AddsTagToGopher(gopherId, tagIds).RunWith(runner).Exec()
	if err != nil {
		log.Panicln("can't add gopher tags: ", err)
	}

	return r.findById(gopherId, runner)
}

// AddUrls creates a link record for each tag id with the given gopherId.
// It panics if can't insert into the db or can't parse the result rows to []Tag.
func (r GopherTagRepository) AddTags(gopherId int64, tagIds []int) []model.Tag {
	return r.addTags(gopherId, tagIds, db.Instance)
}

// AddUrlsWithTx creates a link record for each tag id with the given gopherId.
// Allows to run the query inside a Transaction context.
// It panics if can't insert into the db or can't parse the result rows to []Tag.
func (r GopherTagRepository) AddTagsWithTx(tx *sql.Tx, gopherId int64, tagIds []int) []model.Tag {
	return r.addTags(gopherId, tagIds, tx)
}
