package middleware

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	ctx "github.com/gorilla/context"
	"gitlab.com/patricksangian/go-rest-mux/helpers/utils"
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/auth/model"
)

// VerifyAccessToken middleware
func VerifyAccessToken(next http.HandlerFunc) http.HandlerFunc {
	verifyKey := utils.GetRSAPublicKey()
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")

		if authorizationHeader == "" {
			err := wrapper.Error(http.StatusUnauthorized, "authorization header is required")
			wrapper.Response(res, err.Code, err, err.Message)
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")

		if len(bearerToken) != 2 {
			err := wrapper.Error(http.StatusUnauthorized, "invalid authorization header format")
			wrapper.Response(res, err.Code, err, err.Message)
			return
		}

		tokenString := bearerToken[1]

		token, err := jwt.ParseWithClaims(tokenString, &model.BearerClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return verifyKey, nil
		})

		if err != nil {
			if vE, ok := err.(*jwt.ValidationError); ok {
				var errorString string
				if vE.Errors&jwt.ValidationErrorMalformed != 0 {
					errorString = fmt.Sprintf("invalid token format: %s", tokenString)
				} else if vE.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					errorString = "token is expired"
				} else {
					errorString = fmt.Sprintf("Token Parsing Error: %s", err.Error())
				}
				err := wrapper.Error(http.StatusUnauthorized, errorString)
				wrapper.Response(res, err.Code, err, err.Message)
				return
			}

			err := wrapper.Error(http.StatusUnauthorized, "unknown token error")
			wrapper.Response(res, err.Code, err, err.Message)
			return
		}

		if _, ok := token.Claims.(*model.BearerClaims); token.Valid && ok {
			ctx.Set(req, "decoded", token.Claims)
			next(res, req)
			return
		}

		unknownError := wrapper.Error(http.StatusUnauthorized, "unknown token format")
		wrapper.Response(res, unknownError.Code, unknownError, unknownError.Message)
		return
	})
}
