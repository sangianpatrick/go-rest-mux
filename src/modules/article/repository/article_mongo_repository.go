package repository

import (
	"net/http"

	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article/model"
	"gopkg.in/mgo.v2"
)

type articleMongo struct {
	sess       *mgo.Session
	dbName     string
	collection string
}

// NewArticleMongo acts like a constructor
func NewArticleMongo(mgoSESS *mgo.Session) article.MongoRepository {
	return &articleMongo{
		sess:       mgoSESS,
		dbName:     "gotest",
		collection: "article",
	}
}

func (am *articleMongo) InsertOne(article *model.Article) *wrapper.Property {
	result := make(chan *wrapper.Property)
	go func() {
		err := am.sess.DB(am.dbName).C(am.collection).Insert(article)
		if err != nil {
			result <- wrapper.Error(http.StatusInternalServerError, err.Error())
		}
		result <- wrapper.Data(http.StatusCreated, article, "record has successfuly added")
	}()

	return <-result
}
