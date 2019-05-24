package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/auth"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/auth/model"
)

// authHTTPHandler contains http handler, auth entity and behavior
type authHTTPHandler struct {
	authDomain auth.Domain
}

// NewAuthHTTPHandler act as handler constructor
func NewAuthHTTPHandler(r *mux.Router, ad auth.Domain) {
	ah := &authHTTPHandler{
		authDomain: ad,
	}
	r.HandleFunc("/signin", ah.DoSignIn).Methods("POST")
}

// DoSignIn return user access token
func (ah *authHTTPHandler) DoSignIn(res http.ResponseWriter, req *http.Request) {
	var auth model.Auth
	err := json.NewDecoder(req.Body).Decode(&auth)
	if err != nil {
		data := wrapper.Error(http.StatusUnprocessableEntity, err.Error())
		wrapper.Response(res, data.Code, data, data.Message)
		return
	}
	data := ah.authDomain.SignIn(&auth)
	wrapper.Response(res, data.Code, data, data.Message)
}
