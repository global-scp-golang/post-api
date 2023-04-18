package routes

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"post-api/src/agent/controllers"
)

func SetPostApi(router *mux.Router) *mux.Router {
	// jwt 기능 추가시 decorate 적용
	router.HandleFunc("/api/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/api/posts", controllers.FindPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id:[0-9]+}", controllers.UpdatePosts).Methods("PATCH")
	router.HandleFunc("/api/posts/{id:[0-9]+}", controllers.DeletePost).Methods("DELETE")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return router
}
