package article

import (
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/modules/article/model"
)

// Domain contains article use cases
type Domain interface {
	Create(article *model.Article) *wrapper.Property
}
