package repository

import (
	"database/sql"
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/query"
	"com.github.dazsanchez/gophers-store/scanner"
	sq "github.com/Masterminds/squirrel"
)

type GopherPhotoUrlRepository struct{}

// GopherPhotoUrl allows to manipulate database's gopher_photo_url table.
var GopherPhotoUrl GopherPhotoUrlRepository = GopherPhotoUrlRepository{}

func (r GopherPhotoUrlRepository) findById(id int64, runner sq.BaseRunner) []string {
	uRows, err := query.FindGopherPhotoUrls(id).RunWith(runner).Query()
	if err != nil {
		log.Panicln("can't fetch gopher photo urls: ", err)
	}
	defer uRows.Close()

	urls, err := scanner.ToPhotoUrls(uRows)
	if err != nil {
		log.Panicln("can't parse gopher photo urls: ", err)
	}

	return urls
}

// FindById retrieves all records that matches the given gopherId as []string.
// It panics if can't retrieve the data or can't parse to []string.
func (r GopherPhotoUrlRepository) FindById(id int64) []string {
	return r.findById(id, db.Instance)
}

func (r GopherPhotoUrlRepository) addUrls(gopherId int64, urls []string, runner sq.BaseRunner) []string {
	_, err := query.AddGopherPhotoUrls(gopherId, urls).RunWith(runner).Exec()
	if err != nil {
		log.Panicln("can't add gopher photo urls: ", err)
	}

	return r.findById(gopherId, runner)
}

// AddUrls creates a link record for each url with the given gopherId.
// It panics if can't insert into the db or can't parse the result rows to []string.
func (r GopherPhotoUrlRepository) AddUrls(gopherId int64, urls []string) []string {
	return r.addUrls(gopherId, urls, db.Instance)
}

// AddUrlsWithTx creates a link record for each url with the given gopherId.
// Allows to run the query inside a Transaction context.
// It panics if can't insert into the db or can't parse the result rows to []string.
func (r GopherPhotoUrlRepository) AddUrlsWithTx(tx *sql.Tx, gopherId int64, urls []string) []string {
	return r.addUrls(gopherId, urls, tx)
}
