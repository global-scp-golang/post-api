package routes

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetPostApi(router)
	return router
}
