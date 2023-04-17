package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"post-api/src/agent/common"
	"post-api/src/agent/common/utils"
	"post-api/src/agent/database"
	"post-api/src/agent/services"
	"post-api/src/agent/services/models"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	postService := new(services.PostService)
	postService.DB = database.Connect().DB
	// request -> body 추출
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		panic(err)
	}

	// servicec method 호출
	resource, err := postService.CreatePost(&post)
	//

	if err != nil {
		utils.ResourceWithError(w, common.StatusInternalServerError, err.Error())
	} else {
		utils.ResourceWithJson(w, common.StatusOK, resource)
	}
}
