package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	ctx "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/middleware"
	articleModel "github.com/sangianpatrick/go-rest-mux/src/modules/article/model"
	authModel "github.com/sangianpatrick/go-rest-mux/src/modules/auth/model"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/model"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/validation"
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
	r.HandleFunc("/registration", middleware.VerifyBasicAuth(uh.CreateUser)).Methods("POST")
	r.HandleFunc("/profile/me", middleware.VerifyAccessToken(uh.UpdateProfile)).Methods("PUT")
	r.HandleFunc("/profile/me", middleware.VerifyAccessToken(uh.DeleteProfile)).Methods("DELETE")
	r.HandleFunc("/articles", middleware.VerifyAccessToken(uh.CreateUserArticle)).Methods("POST")
}

// CreateUser will handle creation of user
func (uh *UserHTTPHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		data := wrapper.Error(http.StatusUnprocessableEntity, "unprocessable request payload")
		wrapper.Response(res, data.Code, data, data.Message)
		return
	}
	isValidPayload := validation.IsValidUserRegistrationPayload(&user)
	if !isValidPayload.Success {
		wrapper.Response(res, isValidPayload.Code, isValidPayload, isValidPayload.Message)
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
	var bearer authModel.BearerClaims
	decoded := ctx.Get(req, "decoded")

	mapstructure.Decode(decoded.(*authModel.BearerClaims), &bearer)

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

// UpdateProfile will update user property
func (uh *UserHTTPHandler) UpdateProfile(res http.ResponseWriter, req *http.Request) {
	var updateData model.User
	var bearer authModel.BearerClaims
	decoded := ctx.Get(req, "decoded")

	mapstructure.Decode(decoded.(*authModel.BearerClaims), &bearer)

	err := json.NewDecoder(req.Body).Decode(&updateData)
	if err != nil {
		data := wrapper.Error(http.StatusUnprocessableEntity, "unprocessable payload")
		wrapper.Response(res, data.Code, data, data.Message)
		return
	}
	isValidPayload := validation.IsValidUserUpdateProfilePayload(&updateData)
	if !isValidPayload.Success {
		wrapper.Response(res, isValidPayload.Code, isValidPayload, isValidPayload.Message)
		return
	}
	data := uh.userDomain.Update(bearer.Subject, &updateData)
	wrapper.Response(res, data.Code, data, data.Message)
}

// DeleteProfile remove profile in application
func (uh *UserHTTPHandler) DeleteProfile(res http.ResponseWriter, req *http.Request) {
	var bearer authModel.BearerClaims
	decoded := ctx.Get(req, "decoded")

	mapstructure.Decode(decoded.(*authModel.BearerClaims), &bearer)

	data := uh.userDomain.Delete(bearer.Subject)
	wrapper.Response(res, data.Code, data, data.Message)
}

// CreateUserArticle will create user's article
func (uh *UserHTTPHandler) CreateUserArticle(res http.ResponseWriter, req *http.Request) {
	var bearer authModel.BearerClaims
	var article articleModel.Article

	json.NewDecoder(req.Body).Decode(&article)

	decoded := ctx.Get(req, "decoded")

	mapstructure.Decode(decoded.(*authModel.BearerClaims), &bearer)

	data := uh.userDomain.CreateArticle(bearer.Subject, &article)
	wrapper.Response(res, data.Code, data, "user's article is created")
}
