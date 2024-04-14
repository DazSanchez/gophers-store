package query

import sq "github.com/Masterminds/squirrel"

// FindGopherPhotoUrls selects photo_url field from all records where gopher_id matches the given one.
func FindGopherPhotoUrls(gopherId int64) sq.SelectBuilder {
	return sq.Select("gpu.photo_url").From("gopher_photo_url AS gpu").Where(sq.Eq{
		"gpu.gopher_id": gopherId,
	})
}

// AddGopherPhotoUrls inserts into gopher_photo_url table the list of urls while linking with the given gopher_id.
func AddGopherPhotoUrls(gopherId int64, urls []string) sq.InsertBuilder {
	q := sq.Insert("gopher_photo_url").Columns("gopher_id", "photo_url")

	for _, url := range urls {
		q = q.Values(gopherId, url)
	}

	return q
}
