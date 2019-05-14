package app

import (
	"github.com/gorilla/mux"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/domain"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/handler"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/repository"
	mgo "gopkg.in/mgo.v2"
)

// MountUserApp will run user app
func MountUserApp(prefix string, route *mux.Router, mgoSESS *mgo.Session) {
	mgoRepo := repository.NewUserMongoRepository(mgoSESS)
	domain := domain.NewUserDomain(mgoRepo)
	userRoute := route.PathPrefix(prefix).Subrouter()
	handler.NewUserHTTPHandler(userRoute, domain)
}
