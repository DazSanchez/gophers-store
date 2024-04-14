package query

import sq "github.com/Masterminds/squirrel"

// FindGopherTags retrieves all records from tag table that matches the given gopher_id.
func FindGopherTags(gopherId int64) sq.SelectBuilder {
	return sq.Select("t.id, t.name").From("gopher_tag AS gt").InnerJoin("tag AS t ON gt.tag_id = t.id").Where(sq.Eq{
		"gt.gopher_id": gopherId,
	})
}

// AddsTagToGopher inserts into gopher_tag table all tag ids while linking with the given gopher_id.
func AddsTagToGopher(gopherId int64, tagIds []int) sq.InsertBuilder {
	q := sq.Insert("gopher_tag").Columns("gopher_id", "tag_id")

	for _, tag := range tagIds {
		q = q.Values(gopherId, tag)
	}

	return q
}
