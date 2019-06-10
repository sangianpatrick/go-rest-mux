package app

import (
	"github.com/sangianpatrick/go-rest-mux/src/modules/article"
	"github.com/sangianpatrick/go-rest-mux/src/modules/article/domain"
	"github.com/sangianpatrick/go-rest-mux/src/modules/article/handler"
	"github.com/sangianpatrick/go-rest-mux/src/modules/article/repository"
	mgo "gopkg.in/mgo.v2"
)

// MountArticleEventSource will run article event handler
func MountArticleEventSource(mgoSESS *mgo.Session) article.EventHandler {
	mgoRepo := repository.NewArticleMongo(mgoSESS)
	domain := domain.NewArticleDomain(mgoRepo)
	eventHandler := handler.NewArticleEventHandler(domain)
	return eventHandler
}
