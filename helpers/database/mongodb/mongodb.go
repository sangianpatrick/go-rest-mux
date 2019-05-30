package mongodb

import (
	"os"

	"gitlab.com/patricksangian/go-rest-mux/helpers/logger"
	mgo "gopkg.in/mgo.v2"
)

// NewMongoDBSession returns new sessaion of mongodb.
func NewMongoDBSession() *mgo.Session {
	mongoDBUrl := os.Getenv("MONGO_URL")
	session, err := mgo.Dial(mongoDBUrl)
	if err != nil {
		logger.Fatal("NewMongoDBSession", err)
	}

	err = session.Ping()
	if err != nil {
		logger.Fatal("NewMongoDBSession", err)
	}
	logger.Info("NewMongoDBSession", "Database is connected")

	return session
}
