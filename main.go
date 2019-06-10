package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sangianpatrick/go-rest-mux/middleware"
	"github.com/sangianpatrick/go-rest-mux/src/app"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	mongoDB "github.com/sangianpatrick/go-rest-mux/helpers/database/mongodb"
	"github.com/sangianpatrick/go-rest-mux/helpers/eventemitter"
	"github.com/sangianpatrick/go-rest-mux/helpers/logger"
	"github.com/sangianpatrick/go-rest-mux/helpers/utils"
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/eventsource"
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
	es := eventsource.NewEventSource(MgoSESS)
	emitter := eventemitter.NewEventEmitter(es)

	r := mux.NewRouter()
	r.Use(middleware.SetDefaultHeaders)
	r.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := wrapper.Data(http.StatusOK, nil, "connected to application")
		emitter.EmitPrint(data)
		wrapper.Response(w, data.Code, data, data.Message)
	})

	app.MountAuthApp(r, SignKey, MgoSESS) // Auth
	app.MountUserApp(r, MgoSESS, emitter) // User

	// CORS
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"X-Requested-With", "Origin", "Content-Type", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(r)

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), handler)
	if err != nil {
		logger.Fatal("main.main()", err)
	}
}
