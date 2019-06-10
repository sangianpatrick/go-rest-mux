package article

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article/model"
)

// EventHandler contains article event behavior
type EventHandler interface {
	CreateArticle(article *model.Article) *wrapper.Property
}
