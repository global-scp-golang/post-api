package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"post-api/src/agent/databases"
	"post-api/src/agent/routes"
	"post-api/src/agent/services/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db := databases.OpenDb()
	defer db.Close()
	databases.SetDataBase(db)

	assert := assert.New(t)
	
	m := routes.InitRoutes()

	ts := httptest.NewServer(m)
	defer ts.Close()

	postCreateDto := new(models.Post)
	postCreateDto.Title = "title"
	postCreateDto.Content = "content"
	
	postCreateDtoJson, err := json.Marshal(postCreateDto)
	postCreateDtoByte := bytes.NewBuffer(postCreateDtoJson)
	res, err := http.Post(ts.URL + "/api/posts", "application/json", postCreateDtoByte)

	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
}