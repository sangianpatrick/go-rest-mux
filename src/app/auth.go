package app

import (
	"crypto/rsa"

	"github.com/gorilla/mux"
	"github.com/sangianpatrick/go-rest-mux/src/modules/auth/domain"
	"github.com/sangianpatrick/go-rest-mux/src/modules/auth/handler"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/repository"
	mgo "gopkg.in/mgo.v2"
)

// MountAuthApp will run auth app
func MountAuthApp(route *mux.Router, signKey *rsa.PrivateKey, mgoSESS *mgo.Session) {
	prefixRoute := "/api/v1/auth"
	userMgoRepo := repository.NewUserMongo(mgoSESS)
	domain := domain.NewAuthDomain(signKey, userMgoRepo)
	r := route.PathPrefix(prefixRoute).Subrouter()
	handler.NewAuthHTTPHandler(r, domain)
}
