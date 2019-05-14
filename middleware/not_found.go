package middleware

import (
	"net/http"
	"os"

	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
)

// NotFoundHandler will handle if route is not found
func NotFoundHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Server", os.Getenv("APP_NAME"))
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods",
		"GET,PUT,POST,DELETE")
	res.Header().Set("Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type")
	data := wrapper.Error(http.StatusNotFound, "resource is not found")
	wrapper.Response(res, data.Code, data, data.Message)
}
