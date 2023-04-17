package services

import (
	"database/sql"
	_ "github.com/lib/pq"
	"post-api/src/agent/query"
	"post-api/src/agent/services/models"
)

type PostService struct {
	DB *sql.DB
}

func (s *PostService) CreatePost(post *models.Post) (int, error) {
	q := query.CreatePostQuery()
	//// 쿼리 수행
	prepare, err := s.DB.Prepare(q)
	if err != nil {
		return 0, err
	}

	defer prepare.Close()

	var id int
	err = prepare.QueryRow(post.Title, post.Content).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
