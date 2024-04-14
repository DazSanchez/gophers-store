package query

import (
	"com.github.dazsanchez/gophers-store/model"
	sq "github.com/Masterminds/squirrel"
)

// FindGopherById selects a record from gopher table while also joining with category table by gopher.category_id field.
func FindGopherById(id int64) sq.SelectBuilder {
	return sq.Select("g.id, g.name, g.status, c.id, c.name").From("gopher AS g").InnerJoin("category AS c ON g.category_id = c.id").Where(sq.Eq{
		"g.id": id,
	})
}

// CreateGopher inserts into gopher table the given model.
func CreateGopher(gopher model.Gopher) sq.InsertBuilder {
	return sq.Insert("gopher").Columns("name", "category_id", "status").Values(gopher.Name, gopher.Category.Id, gopher.Status)
}
