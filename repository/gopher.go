package repository

import (
	"log"

	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/model"
	"com.github.dazsanchez/gophers-store/query"
	"com.github.dazsanchez/gophers-store/scanner"
)

type gopherRepository struct{}

var Gopher gopherRepository = gopherRepository{}

func (r gopherRepository) FindById(id int) (g model.Gopher, err error) {
	gStmt, err := db.Instance.Prepare(query.FindGopherById)

	if err != nil {
		log.Panicln("can't fetch gopher: ", err)
		return g, err
	}

	defer gStmt.Close()

	result := gStmt.QueryRow(id)
	g, err = scanner.ToGopher(result)

	if err != nil {
		log.Panicln("can't parse gopher: ", err)
		return g, err
	}

	uStmt, err := db.Instance.Prepare(query.FindGopherPhotoUrls)
	if err != nil {
		log.Panicln("can't fetch gopher photo urls: ", err)
		return g, err
	}

	defer uStmt.Close()

	uRows, err := uStmt.Query(g.Id)
	if err != nil {
		log.Panicln("can't fetch gopher photo urls: ", err)
		return g, err
	}

	urls, err := scanner.ToPhotoUrls(uRows)
	if err != nil {
		log.Panicln("can't parse gopher photo urls: ", err)
		return g, err
	}

	g.PhotoUrls = urls

	tStmt, err := db.Instance.Prepare(query.FindGopherTags)
	tRows, err := tStmt.Query(g.Id)

	if err != nil {
		log.Panicln("can't fetch gopher tags: ", err)
		return g, err
	}

	defer tRows.Close()

	tags, err := scanner.ToTags(tRows)
	if err != nil {
		log.Panicln("can't parse gopher tags: ", err)
		return g, err
	}

	g.Tags = tags

	return g, err
}
