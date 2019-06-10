package handler

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article/model"
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
