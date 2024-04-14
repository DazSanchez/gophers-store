package repository

import (
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/model"
	"com.github.dazsanchez/gophers-store/query"
	"com.github.dazsanchez/gophers-store/scanner"
)

type GopherTagRepository struct{}

// GopherTag allows to manipulate database's gopher_tag table.
var GopherTag GopherTagRepository = GopherTagRepository{}

// FindById retrieves all records that matches the given gopherId as []Tag model.
// It panics if can't retrieve the data or can't parse to []Tag.
func (r GopherTagRepository) FindById(gopherId int64) []model.Tag {
	tRows, err := query.FindGopherTags(gopherId).RunWith(db.Instance).Query()
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

// AddUrls creates a link record for each tag id with the given gopherId.
// It panics if can't insert into the db or can't parse the result rows to []Tag.
func (r GopherTagRepository) AddTags(gopherId int64, tagIds []int) []model.Tag {
	_, err := query.AddsTagToGopher(gopherId, tagIds).RunWith(db.Instance).Exec()
	if err != nil {
		log.Panicln("can't add gopher tags: ", err)
	}

	return r.FindById(gopherId)
}
