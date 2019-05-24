package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	ctx "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/middleware"
	amodel "gitlab.com/patricksangian/go-rest-mux/src/modules/auth/model"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/model"
)

// UserHTTPHandler contains http handler, user entity and behavior
type UserHTTPHandler struct {
	userDomain user.Domain
}

// NewUserHTTPHandler act as handler constructor
func NewUserHTTPHandler(r *mux.Router, ud user.Domain) {
	uh := &UserHTTPHandler{
		userDomain: ud,
	}
	r.HandleFunc("", middleware.VerifyAccessToken(uh.GetAllUser)).Queries("page", "{page}", "size", "{size}").Methods("GET")
	r.HandleFunc("/{userID}", middleware.VerifyAccessToken(uh.GetUserByID)).Methods("GET")
	r.HandleFunc("/profile/me", middleware.VerifyAccessToken(uh.GetProfile)).Methods("GET")
	r.HandleFunc("/registration", uh.CreateUser).Methods("POST")
}

// CreateUser will handle creation of user
func (uh *UserHTTPHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		data := wrapper.Error(http.StatusUnprocessableEntity, err.Error())
		wrapper.Response(res, data.Code, data, data.Message)
		return
	}
	data := uh.userDomain.Create(&user)
	wrapper.Response(res, data.Code, data, data.Message)
}

// GetUserByID return user property by spesific ID with JSON serializer
func (uh *UserHTTPHandler) GetUserByID(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID := params["userID"]
	user := uh.userDomain.GetByID(userID)
	wrapper.Response(res, user.Code, user, user.Message)
}

// GetProfile return authenticated user profile
func (uh *UserHTTPHandler) GetProfile(res http.ResponseWriter, req *http.Request) {
	var bearer amodel.BearerClaims
	decoded := ctx.Get(req, "decoded")

	mapstructure.Decode(decoded.(*amodel.BearerClaims), &bearer)

	data := uh.userDomain.GetByID(bearer.Subject)
	wrapper.Response(res, data.Code, data, "user's profile")
}

// GetAllUser returns list of user
func (uh *UserHTTPHandler) GetAllUser(res http.ResponseWriter, req *http.Request) {
	page, _ := strconv.Atoi(req.FormValue("page"))
	size, _ := strconv.Atoi(req.FormValue("size"))

	data := uh.userDomain.GetAll(page, size)
	wrapper.Response(res, data.Code, data, "list of user")
}
