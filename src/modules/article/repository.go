package article

import (
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/modules/article/model"
)

// MongoRepository conatains mongodb behavior for package article
type MongoRepository interface {
	InsertOne(*model.Article) *wrapper.Property
}
