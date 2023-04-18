package services

import (
	"database/sql"
	"post-api/src/agent/query"
	"post-api/src/agent/services/models"

	_ "github.com/lib/pq"
)

type PostService struct {
	DB          *sql.DB
	updatePosts interface{}
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

func (s *PostService) FindAllPost() ([]*models.Post, error) {
	q := query.FindAllPostQuery()
	posts := []*models.Post{}
	//// 쿼리 수행
	rows, err := s.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.CreatedDt, &post.ModifiedDt)
		posts = append(posts, &post)
	}

	return posts, nil
}

func (s *PostService) UpdatePosts(postId int, req models.PostUpdate) (int, error) {
	updatePosts := query.UpdatePosts()

	prepare, err := s.DB.Prepare(updatePosts)
	if err != nil {
		return 0, err
	}

	var id int

	row := prepare.QueryRow(req.Title, req.Content, postId)
	row.Scan(&id)

	return id, nil
}

func (s *PostService) DeletePost(postId int) (int, error) {
	deletePost := query.DeletePost()

	prepare, err := s.DB.Prepare(deletePost)
	if err != nil {
		return 0, err
	}

	var id int

	row := prepare.QueryRow(postId)
	row.Scan(&id)

	return id, nil
}
