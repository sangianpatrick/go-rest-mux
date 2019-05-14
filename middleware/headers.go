package middleware

import (
	"net/http"
	"os"
)

// SetDefaultHeaders middleware
func SetDefaultHeaders(next http.Handler) http.Handler {
	appName := os.Getenv("APP_NAME")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", appName)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods",
			"GET,PUT,POST,DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Origin, X-Requested-With, Content-Type")
		next.ServeHTTP(w, r)
	})
}
