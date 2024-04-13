package query

const (
	FindGopherById = `
		SELECT g.id, g.name, g.status, c.id, c.name 
		FROM gopher AS g 
			INNER JOIN category AS c ON g.category_id = c.id
		WHERE g.id = ?;
	`

	FindGopherTags = `
		SELECT t.id, t.name
		FROM gopher_tag AS gt
			INNER JOIN tag AS t ON gt.tag_id = t.id
		WHERE gt.gopher_id = ?;
	`

	FindGopherPhotoUrls = `
		SELECT gpu.photo_url
		FROM gopher_photo_url AS gpu
		WHERE gpu.gopher_id = ?;
	`
)
