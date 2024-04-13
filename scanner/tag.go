package scanner

import (
	"database/sql"

	"com.github.dazsanchez/gophers-store/model"
)

func ToTags(rows *sql.Rows) ([]model.Tag, error) {
	var ts []model.Tag

	for rows.Next() {
		var t model.Tag

		err := rows.Scan(&t.Id, &t.Name)

		if err != nil {
			return ts, err
		}

		ts = append(ts, t)
	}

	if ts == nil {
		ts = make([]model.Tag, 0)
	}

	return ts, nil
}
