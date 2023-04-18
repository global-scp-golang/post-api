package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"post-api/src/agent/databases"
	"post-api/src/agent/routes"
	"post-api/src/agent/services/models"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ts *httptest.Server

// Before 제일 처음시작할 때 한번?
func TestMain(m *testing.M) {
	db := databases.OpenTestDB()
	defer db.Close()

	databases.SetDataBase(db)

	route := routes.InitRoutes()

	ts = httptest.NewServer(route)
	defer ts.Close()

	m.Run()
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	postCreateDto := new(models.Post)
	postCreateDto.Title = "tx"
	postCreateDto.Content = "tx"

	postCreateDtoJson, err := json.Marshal(postCreateDto)
	postCreateDtoByte := bytes.NewBuffer(postCreateDtoJson)
	res, err := http.Post(ts.URL+"/api/posts", "application/json", postCreateDtoByte)

	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
}

func TestFindAll(t *testing.T) {
	assert := assert.New(t)

	res, err := http.Get(ts.URL + "/api/posts")

	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
}

func TestUpdatePosts(t *testing.T) {
	assert := assert.New(t)

	postId := 2

	postUpdateDto := new(models.PostUpdate)
	postUpdateDto.Title = "update"
	postUpdateDto.Content = "update"

	postCreateDtoJson, err := json.Marshal(postUpdateDto)
	postCreateDtoByte := bytes.NewBuffer(postCreateDtoJson)

	req, err := http.NewRequest("PATCH", ts.URL+"/api/posts/"+strconv.Itoa(postId), postCreateDtoByte)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
}

func TestDeletePosts(t *testing.T) {
	assert := assert.New(t)

	postId := 2

	req, err := http.NewRequest("DELETE", ts.URL+"/api/posts/"+strconv.Itoa(postId), nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
}
