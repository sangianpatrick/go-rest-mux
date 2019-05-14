package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user"
)

// UserHTTPHandler contains http handler, user entity and behavior of users
type UserHTTPHandler struct {
	userDomain user.Domain
}

// NewUserHTTPHandler act as handler constructor
func NewUserHTTPHandler(r *mux.Router, ud user.Domain) {
	uh := &UserHTTPHandler{
		userDomain: ud,
	}
	r.HandleFunc("/{userID}", uh.GetUserByID).Methods("GET")
}

// GetUserByID return user property by spesific ID with JSON serializer
func (uh *UserHTTPHandler) GetUserByID(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID := params["userID"]
	user := uh.userDomain.GetByID(userID)
	wrapper.Response(res, user.Code, user, user.Message)
}
