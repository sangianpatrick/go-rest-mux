package handler

import (
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/modules/article"
	"github.com/sangianpatrick/go-rest-mux/src/modules/article/model"
)

// ArticleEventHandler contains pro
type articleEventHandler struct {
	ad article.Domain
}

// NewArticleEventHandler acts like a constructor
func NewArticleEventHandler(ad article.Domain) article.EventHandler {
	return &articleEventHandler{
		ad: ad,
	}
}

func (ae *articleEventHandler) CreateArticle(article *model.Article) *wrapper.Property {
	result := ae.ad.Create(article)
	return result
}
