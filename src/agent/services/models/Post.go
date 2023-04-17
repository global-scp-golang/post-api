package models

import "time"

type Post struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedDt  time.Time `json:"createdDt"`
	modifiedDt time.Time `json:"modifiedDt"`
}
