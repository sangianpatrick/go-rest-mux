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

	r := mux.NewRouter()
	r.Use(middleware.SetHeaders)
	r.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := wrapper.Data(http.StatusOK, nil, "connected to application")
		wrapper.Response(w, data.Code, data, data.Message)
	})

	app.MountUserApp("/api/v1/user", r, MgoSESS)

	// CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	credsOk := handlers.AllowCredentials()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk, credsOk)(r)))
}
