package scanner

import (
	"database/sql"

	"com.github.dazsanchez/gophers-store/model"
	sq "github.com/Masterminds/squirrel"
)

// ToGopher is a serializer to convert sql.Rows to a Gopher model.
// It panics is there's an error while scanning.
func ToGopher(row sq.RowScanner) (model.Gopher, error) {
	var c model.Category
	var g model.Gopher

	err := row.Scan(&g.Id, &g.Name, &g.Status, &c.Id, &c.Name)

	if err != nil {
		return g, err
	}

	g.Category = c

	return g, nil
}

// ToPhotoUrls is a serializer to convert sql.Rows to a slice of strings.
// It panics is there's an error while scanning.
func ToPhotoUrls(rows *sql.Rows) ([]string, error) {
	us := make([]string, 0)

	for rows.Next() {
		var u string

		err := rows.Scan(&u)

		if err != nil {
			return us, err
		}

		us = append(us, u)
	}

	return us, nil
}
