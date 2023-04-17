package query

const (
	GetPostsQuery = "select"
)

func CreatePostQuery() string {
	query := `INSERT INTO "Posts" (title, content, "createdDt", "modifiedDt") VALUES ($1, $2, current_timestamp, current_timestamp) RETURNING id`
	return query
}
