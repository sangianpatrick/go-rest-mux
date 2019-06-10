package app

import (
	"github.com/gorilla/mux"
	"github.com/sangianpatrick/go-rest-mux/helpers/eventemitter"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/domain"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/handler"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/repository"
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
