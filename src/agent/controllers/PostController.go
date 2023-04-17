package controllers

import (
	"net/http"
	"studySwagger/src/agent/common/utils"
	"studySwagger/src/agent/database"
	"studySwagger/src/agent/services"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	postService := new(services.PostService)
	postService.DB = database.Connect().DB
	// request -> body 추출

	// servicec method 호출
	resource, err := postService.CreatePost()
	//
	if err != nil {
		utils.ResourceWithError(w, errorCode, err)
	} else {
		utils.ResourceWithJson(w, statusCode, resource)
	}
}
