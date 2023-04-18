package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"post-api/src/agent/common"
	"post-api/src/agent/common/utils"
	"post-api/src/agent/databases"
	"post-api/src/agent/services"
	"post-api/src/agent/services/models"
	"strconv"
)

// @Summary Post 생성 APIasasas
// @Router /api/posts [post]
// @Param request body models.Post true "request"
func CreatePost(w http.ResponseWriter, r *http.Request) {
	postService := new(services.PostService)
	postService.DB = databases.Connect().DB
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

// @Summary Post FindAll API
// @Router /api/posts [get]
func FindPosts(w http.ResponseWriter, r *http.Request) {
	postService := new(services.PostService)
	postService.DB = databases.Connect().DB
	// request -> body 추출

	// servicec method 호출
	posts, err := postService.FindAllPost()
	//
	if err != nil {
		utils.ResourceWithError(w, common.StatusInternalServerError, err.Error())
	} else {
		utils.ResourceWithJson(w, common.StatusOK, posts)
	}
}

// @Summary Post Update API
// @Param postId path int true "postId"
// @Router /api/posts/{postId} [patch]
// @Param request body models.PostUpdate true "request"
func UpdatePosts(w http.ResponseWriter, r *http.Request) {
	postService := new(services.PostService)
	postService.DB = databases.Connect().DB
	// request -> body 추출
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ResourceWithError(w, common.StatusInternalServerError, err.Error())
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var req models.PostUpdate
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	postId, err := postService.UpdatePosts(id, req)

	if err != nil {
		utils.ResourceWithError(w, common.StatusInternalServerError, err.Error())
	} else {
		utils.ResourceWithJson(w, common.StatusOK, postId)
	}
}

// @Summary Post Delete API
// @Param postId path int true "postId"
// @Router /api/posts/{postId} [delete]
func DeletePost(w http.ResponseWriter, r *http.Request) {
	postService := new(services.PostService)
	postService.DB = databases.Connect().DB
	// request -> body 추출
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ResourceWithError(w, common.StatusInternalServerError, err.Error())
	}

	postId, err := postService.DeletePost(id)

	if err != nil {
		utils.ResourceWithError(w, common.StatusInternalServerError, err.Error())
	} else {
		utils.ResourceWithJson(w, common.StatusOK, postId)
	}
}

