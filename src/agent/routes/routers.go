package routes

import (
	"github.com/gorilla/mux"
	"studySwagger/src/agent/controllers"
)

func SetPostApi(router *mux.Router) *mux.Router {
	// jwt 기능 추가시 decorate 적용
	router.HandleFunc("/api/posts", controllers.CreatePost).Methods("POST")
	return router
}
