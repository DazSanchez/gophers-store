package query

import sq "github.com/Masterminds/squirrel"

var (
	// FindAllTags select all records from tag table.
	FindAllTags sq.SelectBuilder = sq.Select("*").From("tag")
)
