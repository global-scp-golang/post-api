package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"post-api/src/agent/databases"
	_ "post-api/src/agent/docs"
	"post-api/src/agent/routes"
)

func main() {
	db := databases.OpenDB()
	defer db.Close()
	databases.SetDataBase(db)
	router := routes.InitRoutes()

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// router 서버 실행
	fmt.Println("Server On :8000")
	http.ListenAndServe(":8000", cors(router))
}

