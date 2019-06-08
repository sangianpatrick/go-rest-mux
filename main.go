package main

import (
	"net/http"

	"gitlab.com/patricksangian/go-rest-mux/middleware"
	"gitlab.com/patricksangian/go-rest-mux/src/app"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	mongoDB "gitlab.com/patricksangian/go-rest-mux/helpers/database/mongodb"
	"gitlab.com/patricksangian/go-rest-mux/helpers/eventemitter"
	"gitlab.com/patricksangian/go-rest-mux/helpers/logger"
	"gitlab.com/patricksangian/go-rest-mux/helpers/utils"
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
)

func init() {
	err := godotenv.Load(`.env`)
	if err != nil {
		logger.Fatal("main.init()", err)
	}
}

func main() {
	MgoSESS := mongoDB.NewMongoDBSession()
	SignKey := utils.GetRSAPrivateKey()
	emitter := eventemitter.NewEventEmitter()

	r := mux.NewRouter()
	r.Use(middleware.SetDefaultHeaders)
	r.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := wrapper.Data(http.StatusOK, nil, "connected to application")
		emitter.EmitPrint(data)
		wrapper.Response(w, data.Code, data, data.Message)
	})

	app.MountAuthApp(r, SignKey, MgoSESS) // Auth
	app.MountUserApp(r, MgoSESS)          // User

	// CORS
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"X-Requested-With", "Origin", "Content-Type", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(r)

	err := http.ListenAndServe("localhost:9000", handler)
	if err != nil {
		logger.Fatal("main.main()", err)
	}
}
