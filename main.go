package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/patricksangian/go-rest-mux/middleware"
	"gitlab.com/patricksangian/go-rest-mux/src/app"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	db "gitlab.com/patricksangian/go-rest-mux/helpers/database"
	"gitlab.com/patricksangian/go-rest-mux/helpers/utils"
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
)

func init() {
	err := godotenv.Load(`.env`)
	if err != nil {
		panic(err)
	}
}

func main() {
	MgoSESS := db.NewMongoDBSession()
	SignKey := utils.GetRSAPrivateKey()

	r := mux.NewRouter()
	r.Use(middleware.SetDefaultHeaders)
	r.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := wrapper.Data(http.StatusOK, nil, "connected to application")
		wrapper.Response(w, data.Code, data, data.Message)
	})

	app.MountAuthApp(r, SignKey, MgoSESS) // Auth
	app.MountUserApp(r, MgoSESS)          // Users

	// CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	credsOk := handlers.AllowCredentials()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk, credsOk)(r)))
}
