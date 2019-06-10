package app

import (
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article/domain"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article/handler"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article/repository"
	mgo "gopkg.in/mgo.v2"
)

// MountArticleEventSource will run article event handler
func MountArticleEventSource(mgoSESS *mgo.Session) article.EventHandler {
	mgoRepo := repository.NewArticleMongo(mgoSESS)
	domain := domain.NewArticleDomain(mgoRepo)
	eventHandler := handler.NewArticleEventHandler(domain)
	return eventHandler
}
