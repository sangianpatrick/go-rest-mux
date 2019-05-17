package app

import (
	"crypto/rsa"

	"github.com/gorilla/mux"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/auth/domain"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/auth/handler"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/repository"
	mgo "gopkg.in/mgo.v2"
)

// MountAuthApp will run auth app
func MountAuthApp(route *mux.Router, signKey *rsa.PrivateKey, mgoSESS *mgo.Session) {
	prefixRoute := "/api/v1/auth"
	userMgoRepo := repository.NewUserMongoRepository(mgoSESS)
	domain := domain.NewAuthDomain(signKey, userMgoRepo)
	userRoute := route.PathPrefix(prefixRoute).Subrouter()
	handler.NewAuthHTTPHandler(userRoute, domain)
}
