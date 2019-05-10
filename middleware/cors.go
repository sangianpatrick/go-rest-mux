package middleware

import (
	"net/http"
	"os"

	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
)

// CORS middleware
func CORS(next http.Handler) http.Handler {
	appName := os.Getenv("APP_NAME")
	acceptMethod := []string{"POST", "GET", "PUT", "DELETE"}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		w.Header().Set("Server", appName)
		w.Header().Set("Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods",
			"GET,PUT,POST,DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Origin, X-Requested-With, Content-Type, Accept")
		if method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Write(nil)
			// return
		} else {
			methodMatch := func() bool {
				for _, v := range acceptMethod {
					if method == v {
						return true
					}
				}
				return false
			}()
			if !methodMatch {
				err := wrapper.Error(http.StatusMethodNotAllowed, "Method is not allowed")
				wrapper.Response(w, err.Code, &err, err.Message)
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
