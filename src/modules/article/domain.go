package article

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/article/model"
)

// Domain contains article use cases
type Domain interface {
	Create(article *model.Article) *wrapper.Property
}
