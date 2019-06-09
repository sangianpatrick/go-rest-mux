package app

import (
	"github.com/gorilla/mux"
	"gitlab.com/patricksangian/go-rest-mux/helpers/eventemitter"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/domain"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/handler"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/repository"
	mgo "gopkg.in/mgo.v2"
)

// MountUserApp will run user app
func MountUserApp(route *mux.Router, mgoSESS *mgo.Session, emitter eventemitter.Emitter) {
	prefixRoute := "/api/v1/users"
	mgoRepo := repository.NewUserMongo(mgoSESS)
	domain := domain.NewUserDomain(mgoRepo, emitter)
	r := route.PathPrefix(prefixRoute).Subrouter()
	handler.NewUserHTTPHandler(r, domain)
}
