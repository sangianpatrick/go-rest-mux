package repository

import (
	"net/http"

	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserMongoRepository contains property and bevaior of mongodb in order to fetch and/or retrive user data from mongodb.
type UserMongoRepository struct {
	sess       *mgo.Session
	dbName     string
	collection string
}

// NewUserMongoRepository acts as constructor.
func NewUserMongoRepository(mgoSESS *mgo.Session) user.MongoRepositrory {
	return &UserMongoRepository{
		sess:       mgoSESS,
		dbName:     "gotest",
		collection: "user",
	}
}

// InsertOne will add new record to user collection
func (umr *UserMongoRepository) InsertOne(user *model.User) *wrapper.Property {
	result := make(chan *wrapper.Property)
	next := make(chan bool)

	go func() {
		data := umr.FindByEmail(user.Email)
		if !data.Error {
			result <- wrapper.Error(http.StatusConflict, "duplicate data entry")
			next <- false
		}
		next <- true
	}()
	go func(n <-chan bool) {
		if <-n {
			err := umr.sess.DB(umr.dbName).C(umr.collection).Insert(user)
			if err != nil {
				result <- wrapper.Error(http.StatusInternalServerError, err.Error())
			}
			result <- wrapper.Data(http.StatusCreated, user, "record has successfuly added")
		}
	}(next)
	return <-result
}

// FindByID returns user data with spesific ID
func (umr *UserMongoRepository) FindByID(ID string) *wrapper.Property {
	var user model.User
	query := bson.M{
		"id": ID,
	}
	projection := bson.M{
		"password": 0,
	}
	result := make(chan *wrapper.Property)

	go func() {
		err := umr.sess.DB(umr.dbName).C(umr.collection).Find(query).Select(projection).One(&user)
		if err != nil {
			result <- wrapper.Error(http.StatusNotFound, err.Error())
		}
		result <- wrapper.Data(http.StatusOK, user, "detail of record")
	}()

	return <-result
}

// FindByEmail returns user data with spesific Email
func (umr *UserMongoRepository) FindByEmail(email string) *wrapper.Property {
	var user model.User
	query := bson.M{
		"email": email,
	}
	result := make(chan *wrapper.Property)

	go func() {
		err := umr.sess.DB(umr.dbName).C(umr.collection).Find(query).One(&user)
		if err != nil {
			result <- wrapper.Error(http.StatusNotFound, err.Error())
		}
		result <- wrapper.Data(http.StatusOK, user, "detail of record")
	}()

	return <-result
}

// FindAll returns list of user.
func (umr *UserMongoRepository) FindAll(limit int, skip int) *wrapper.Property {
	var users model.User
	projection := bson.M{
		"password": 0,
	}
	result := make(chan *wrapper.Property)

	go func() {
		err := umr.sess.DB(umr.dbName).C(umr.collection).Find(nil).Select(projection).Limit(limit).Skip(skip).All(&users)
		if err != nil {
			result <- wrapper.Error(http.StatusNotFound, err.Error())
		}
		result <- wrapper.Data(http.StatusOK, users, "list of user")
	}()

	return <-result
}
