package services

import (
	"database/sql"
)

type PostService struct {
	DB *sql.DB
}

func (s *PostService) CreatePost() {
	//q := query.CreatePostQuery
	//// 쿼리 수행
	//s.DB.Exec(q)
	//
	//// 수행 결과 객체로 만들기
	//resource := []models.Post{}
	//
	//return resource, err
}
