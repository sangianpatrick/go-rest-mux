package article

import (
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/modules/article/model"
)

// EventHandler contains article event behavior
type EventHandler interface {
	CreateArticle(article *model.Article) *wrapper.Property
}
