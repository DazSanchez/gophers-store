package query

import sq "github.com/Masterminds/squirrel"

var (
	// FindAllCategories selects all records from category table.
	FindAllCategories sq.SelectBuilder = sq.Select("*").From("category")
)
