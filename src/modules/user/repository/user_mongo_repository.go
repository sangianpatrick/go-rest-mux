package repository

import (
	"math"
	"net/http"

	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/model"
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

func (um *userMongo) customFetch(fn func()) {
	go fn()
}

// InsertOne will add new record to user collection
func (um *userMongo) InsertOne(user *model.User) *wrapper.Property {
	result := make(chan *wrapper.Property)
	go func() {
		err := um.sess.DB(um.dbName).C(um.collection).Insert(user)
		if err != nil {
			if mgo.IsDup(err) {
				result <- wrapper.Error(http.StatusConflict, "user is already exist")
			}
			result <- wrapper.Error(http.StatusInternalServerError, err.Error())
		}
		result <- wrapper.Data(http.StatusCreated, user, "record has successfuly added")
	}()
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
			if err == mgo.ErrNotFound {
				result <- wrapper.Error(http.StatusNotFound, err.Error())
			}
			result <- wrapper.Error(http.StatusInternalServerError, err.Error())
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
			if err == mgo.ErrNotFound {
				result <- wrapper.Error(http.StatusNotFound, err.Error())
			}
			result <- wrapper.Error(http.StatusInternalServerError, err.Error())
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
				errChan <- wrapper.Error(http.StatusInternalServerError, err.Error())
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

func (um *userMongo) DeleteOne(ID string) *wrapper.Property {
	result := make(chan *wrapper.Property)
	go func() {
		err := um.sess.DB(um.dbName).C(um.collection).Remove(bson.M{"id": ID})
		if err != nil {
			if err == mgo.ErrNotFound {
				result <- wrapper.Error(http.StatusNotFound, err.Error())
			}
			result <- wrapper.Error(http.StatusInternalServerError, err.Error())
		}
		result <- wrapper.Data(http.StatusOK, nil, "user is successfuly deleted")
	}()

	return <-result
}

func (um *userMongo) UpdateOne(ID string, data *model.User) *wrapper.Property {
	var userBson bson.M
	userBsonByte, _ := bson.Marshal(data)
	bson.Unmarshal(userBsonByte, &userBson)
	updateData := bson.M{"$set": userBson}
	result := make(chan *wrapper.Property)
	go func() {
		err := um.sess.DB(um.dbName).C(um.collection).Update(bson.M{"id": ID}, updateData)
		if err != nil {
			if err == mgo.ErrNotFound {
				result <- wrapper.Error(http.StatusNotFound, err.Error())
			}
			result <- wrapper.Error(http.StatusInternalServerError, err.Error())
		}
		result <- wrapper.Data(http.StatusOK, nil, "user is succesfuly updated")
	}()
	return <-result
}
