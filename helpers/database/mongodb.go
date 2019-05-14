package database

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

// NewMongoDBSession returns new sessaion of mongodb.
func NewMongoDBSession() *mgo.Session {
	host := os.Getenv("MONGO_HOST")
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	dbname := os.Getenv("MONGO_DBNAME")
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{host},
		Username: user,
		Password: password,
		Database: dbname,
	})
	if err != nil {
		log.Fatalf(`MongoDB Error: %s`, err)
	}

	err = session.Ping()
	if err != nil {
		log.Fatalf(`MongoDB Error: %s`, err)
	}
	log.Println("MongoDB Info: Database is connected")

	return session
}
