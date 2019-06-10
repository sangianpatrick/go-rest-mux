package middleware

import (
	"net/http"
	"os"

	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
)

// VerifyBasicAuth middleware
func VerifyBasicAuth(next http.HandlerFunc) http.HandlerFunc {
	username := os.Getenv("BASIC_AUTH_USERNAME")
	password := os.Getenv("BASIC_AUTH_PASSWORD")
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var errData *wrapper.Property
		u, p, ok := req.BasicAuth()
		if !ok {
			errData = wrapper.Error(http.StatusForbidden, "unknown request")
		}
		if !(u == username && p == password) {
			errData = wrapper.Error(http.StatusUnauthorized, "unauthorized request")
		}
		if errData != nil {
			wrapper.Response(res, errData.Code, errData, errData.Message)
			return
		}
		next(res, req)
	})
}
