package repository

import (
	"math"
	"net/http"

	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// userMongoRepository contains property and bevaior of mongodb in order to fetch and/or retrive user data from mongodb.
type userMongo struct {
	sess       *mgo.Session
	dbName     string
	collection string
}

// NewUserMongo acts as constructor.
func NewUserMongo(mgoSESS *mgo.Session) user.MongoRepositrory {
	return &userMongo{
		sess:       mgoSESS,
		dbName:     "gotest",
		collection: "user",
	}
}

// InsertOne will add new record to user collection
func (um *userMongo) InsertOne(user *model.User) *wrapper.Property {
	result := make(chan *wrapper.Property)
	next := make(chan bool)

	go func() {
		data := um.FindByEmail(user.Email)
		if data.Success {
			result <- wrapper.Error(http.StatusConflict, "user email is already registered")
			next <- false
		}
		next <- true
	}()
	go func(n <-chan bool) {
		if <-n {
			err := um.sess.DB(um.dbName).C(um.collection).Insert(user)
			if err != nil {
				result <- wrapper.Error(http.StatusInternalServerError, err.Error())
			}
			result <- wrapper.Data(http.StatusCreated, user, "record has successfuly added")
		}
	}(next)
	return <-result
}

// FindByID returns user data with spesific ID
func (um *userMongo) FindByID(ID string) *wrapper.Property {
	var user model.User
	query := bson.M{
		"id": ID,
	}
	projection := bson.M{
		"password": 0,
	}
	result := make(chan *wrapper.Property)

	go func() {
		err := um.sess.DB(um.dbName).C(um.collection).Find(query).Select(projection).One(&user)
		if err != nil {
			result <- wrapper.Error(http.StatusNotFound, err.Error())
		}
		result <- wrapper.Data(http.StatusOK, user, "detail of record")
	}()

	return <-result
}

// FindByEmail returns user data with spesific Email
func (um *userMongo) FindByEmail(email string) *wrapper.Property {
	var user model.User
	query := bson.M{
		"email": email,
	}
	result := make(chan *wrapper.Property)

	go func() {
		err := um.sess.DB(um.dbName).C(um.collection).Find(query).One(&user)
		if err != nil {
			result <- wrapper.Error(http.StatusNotFound, err.Error())
		}
		result <- wrapper.Data(http.StatusOK, user, "detail of record")
	}()

	return <-result
}

// FindAll returns list of user.
func (um *userMongo) FindAll(page int, size int) *wrapper.Property {
	var users model.Users
	errChan := make(chan *wrapper.Property)
	totalData := make(chan int)
	data := make(chan model.Users)
	result := make(chan *wrapper.Property)
	skip := (page - 1) * size
	projection := bson.M{
		"password": 0,
	}

	go func() {
		go func() {
			countAll, err := um.sess.DB(um.dbName).C(um.collection).Count()
			if err != nil {
				errChan <- wrapper.Error(http.StatusNotFound, err.Error())
			}
			totalData <- countAll
		}()

		go func() {
			err := um.sess.DB(um.dbName).C(um.collection).Find(nil).Select(projection).Limit(size).Skip(skip).All(&users)
			if err != nil {
				errChan <- wrapper.Error(http.StatusNotFound, err.Error())
			}
			if len(users) < 1 {
				errChan <- wrapper.Error(http.StatusNotFound, "page is empty")
			}
			data <- users
		}()

		td := <-totalData
		d := <-data
		tp := int(math.Ceil(float64(td) / float64(size)))
		tdp := len(d)
		result <- wrapper.PaginationData(http.StatusOK, d, tp, page, td, tdp, "list of user")
	}()

	select {
	case err := <-errChan:
		return err
	case res := <-result:
		return res
	}
}
