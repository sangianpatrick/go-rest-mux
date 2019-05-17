package app

import (
	"github.com/gorilla/mux"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/domain"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/handler"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/repository"
	mgo "gopkg.in/mgo.v2"
)

// MountUserApp will run user app
func MountUserApp(route *mux.Router, mgoSESS *mgo.Session) {
	prefixRoute := "/api/v1/user"
	mgoRepo := repository.NewUserMongoRepository(mgoSESS)
	domain := domain.NewUserDomain(mgoRepo)
	userRoute := route.PathPrefix(prefixRoute).Subrouter()
	handler.NewUserHTTPHandler(userRoute, domain)
}
