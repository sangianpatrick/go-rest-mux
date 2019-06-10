package model

import (
	uModel "github.com/sangianpatrick/go-rest-mux/src/modules/user/model"
)

// Article contains article property
type Article struct {
	ID        string      `json:"id"`
	Category  string      `json:"category"`
	Topic     string      `json:"topic"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	CreatedBy uModel.User `json:"createdBy"`
}
