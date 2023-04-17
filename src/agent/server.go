package main

import (
	"fmt"
	"net/http"
	"post-api/src/agent/databases"
	"post-api/src/agent/routes"
)


func main() {
	db := databases.OpenDb()
	defer db.Close()
	databases.SetDataBase(db)
	router := routes.InitRoutes()

	// router 서버 실행
	fmt.Println("Server On :8000")
	http.ListenAndServe(":8000", router)
}
