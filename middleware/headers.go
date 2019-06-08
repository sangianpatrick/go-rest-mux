package middleware

import (
	"net/http"
	"os"
)

// SetDefaultHeaders middleware
func SetDefaultHeaders(next http.Handler) http.Handler {
	appName := os.Getenv("APP_NAME")
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Server", appName)
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods",
			"GET, PUT, POST, DELETE")
		res.Header().Set("Access-Control-Allow-Headers",
			"Origin, X-Requested-With, Content-Type")
		next.ServeHTTP(res, req)
	})
}
