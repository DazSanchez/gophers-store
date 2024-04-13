package scanner

import (
	"database/sql"

	"com.github.dazsanchez/gophers-store/model"
)

func ToGopher(row *sql.Row) (model.Gopher, error) {
	var c model.Category
	var g model.Gopher

	err := row.Scan(&g.Id, &g.Name, &g.Status, &c.Id, &c.Name)

	if err != nil {
		return g, err
	}

	g.Category = c

	return g, nil
}

func ToPhotoUrls(rows *sql.Rows) ([]string, error) {
	var us []string

	for rows.Next() {
		var u string

		err := rows.Scan(&u)

		if err != nil {
			return us, err
		}

		us = append(us, u)
	}

	if us == nil {
		us = make([]string, 0)
	}

	return us, nil
}
