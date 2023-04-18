package query

const (
	GetPostsQuery = "select"
)

func CreatePostQuery() string {
	query := `INSERT INTO "Posts" (title, content, "createdDt", "modifiedDt") VALUES ($1, $2, current_timestamp, current_timestamp) RETURNING id`
	return query
}

func FindAllPostQuery() string {
	query := `SELECT * FROM "Posts"`
	return query
}

func DeleteAllPosts() string {
	query := `DELETE * FROM "Posts"`
	return query
}

func UpdatePosts() string {
	query := `UPDATE "Posts" SET title = $1, content = $2, "modifiedDt" = current_timestamp WHERE id = $3 RETURNING id`
	return query
}

func DeletePost() string {
	query := `DELETE FROM "Posts" WHERE id = $1 RETURNING id`
	return query
}
