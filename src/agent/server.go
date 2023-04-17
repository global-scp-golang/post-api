package main

import (
	"database/sql"
	"time"
)

func main() {
	db, err := sql.Open("", "")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(20) // DB Connection을 사용하지 않을때에도 Connection 풀에 최대 Connection 20개를 둔다
	db.SetMaxOpenConns(10) // DB Connection을 최대 10개까지 동시 사용가능하다
	defer db.Close()
	database.SetDataBase(db)
	router := routes.InitRoutes()

	// router 서버 실행
}
