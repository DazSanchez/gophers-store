package repository

import (
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/query"
	"com.github.dazsanchez/gophers-store/scanner"
)

type GopherPhotoUrlRepository struct{}

// GopherPhotoUrl allows to manipulate database's gopher_photo_url table.
var GopherPhotoUrl GopherPhotoUrlRepository = GopherPhotoUrlRepository{}

// FindById retrieves all records that matches the given gopherId as []string.
// It panics if can't retrieve the data or can't parse to []string.
func (r GopherPhotoUrlRepository) FindById(id int64) []string {
	uRows, err := query.FindGopherPhotoUrls(id).RunWith(db.Instance).Query()
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

// AddUrls creates a link record for each url with the given gopherId.
// It panics if can't insert into the db or can't parse the result rows to []string.
func (r GopherPhotoUrlRepository) AddUrls(gopherId int64, urls []string) []string {
	_, err := query.AddGopherPhotoUrls(gopherId, urls).RunWith(db.Instance).Exec()
	if err != nil {
		log.Panicln("can't add gopher photo urls: ", err)
	}

	return r.FindById(gopherId)
}
