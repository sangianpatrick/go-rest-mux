package domain

import (
	"crypto/rsa"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sangianpatrick/go-rest-mux/helpers/logger"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sangianpatrick/go-rest-mux/helpers/utils"
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/modules/auth"
	"github.com/sangianpatrick/go-rest-mux/src/modules/auth/model"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user"
	userModel "github.com/sangianpatrick/go-rest-mux/src/modules/user/model"
)

// authDomain contains auth property, entity and use cases
type authDomain struct {
	mgoRepo user.MongoRepositrory
	signKey *rsa.PrivateKey
}

// NewAuthDomain act as constructor
func NewAuthDomain(signKey *rsa.PrivateKey, mgoRepo user.MongoRepositrory) auth.Domain {
	return &authDomain{
		mgoRepo: mgoRepo,
		signKey: signKey,
	}
}

func (ad *authDomain) generateCredential(user userModel.User) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	claims := model.BearerClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   user.ID,
			Issuer:    os.Getenv("APP_NAME"),
		},
	}
	t.Claims = claims
	tokenString, err := t.SignedString(ad.signKey)
	if err != nil {
		logger.Fatal("auth_domain.generateCredential", err)
		return "", errors.New("Error while signing token")
	}
	return tokenString, nil
}

// SignIn will return user auth token
func (ad *authDomain) SignIn(payload *model.Auth) *wrapper.Property {
	var credential model.Credential
	retrieve := ad.mgoRepo.FindByEmail(payload.Email)
	if !retrieve.Success {
		return wrapper.Error(http.StatusUnauthorized, "not a registered user")
	}
	user, ok := retrieve.Data.(userModel.User)
	if !ok {
		log.Fatal("Auth Error: Assertion on userModel.User")
		return wrapper.Error(http.StatusInternalServerError, "error detected due to user signin")
	}

	password, err := utils.Decrypt([]byte(utils.SecretKey), user.Password)
	if err != nil {
		log.Fatal("Auth Error: Decrypting user password")
		return wrapper.Error(http.StatusInternalServerError, "error detected due to user signin")
	}
	if payload.Password != password {
		return wrapper.Error(http.StatusBadRequest, "invalid email or password")
	}

	token, err := ad.generateCredential(user)
	if err != nil {
		return wrapper.Error(http.StatusInternalServerError, err.Error())
	}

	credential.Name = user.Name
	credential.AccessToken = token

	return wrapper.Data(http.StatusOK, credential, "signin success")
}
