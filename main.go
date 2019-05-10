package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	db "gitlab.com/patricksangian/go-rest-mux/helpers/database"
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/middleware"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/model"
)

func init() {
	err := godotenv.Load(`.env`)
	if err != nil {
		panic(err)
	}
}

func main() {
	MgoSess := db.NewMongoDBSession()
	err := MgoSess.Ping()
	if err != nil {
		fmt.Printf(`%s`, err)
	}
	r := mux.NewRouter()
	r.Use(middleware.CORS)

	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{
			ID:    "0001",
			Name:  "Patrick Maurits Sangian",
			Email: "patricksangian@gmail.com",
			Phone: "08124541588",
		}
		data := wrapper.Data(http.StatusOK, user, "user data")
		wrapper.Response(w, data.Code, &data, data.Message)
	})

	server := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(`127.0.0.1:%s`, os.Getenv("PORT")),
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
