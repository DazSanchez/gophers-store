package scanner

import (
	"database/sql"

	"com.github.dazsanchez/gophers-store/model"
)

func ToCategories(rows *sql.Rows) ([]model.Category, error) {
	var cs []model.Category

	for rows.Next() {
		var c model.Category

		err := rows.Scan(&c.Id, &c.Name)

		if err != nil {
			return cs, err
		}

		cs = append(cs, c)
	}

	return cs, nil
}
