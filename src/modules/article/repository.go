package article

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article/model"
)

// MongoRepository conatains mongodb behavior for package article
type MongoRepository interface {
	InsertOne(*model.Article) *wrapper.Property
}
