package domain

import (
	"github.com/google/uuid"
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/modules/article"
	"github.com/sangianpatrick/go-rest-mux/src/modules/article/model"
)

type articleDomain struct {
	mgoRepo article.MongoRepository
}

// NewArticleDomain acts like constructor
func NewArticleDomain(mgoRepo article.MongoRepository) article.Domain {
	return &articleDomain{
		mgoRepo: mgoRepo,
	}
}

func (ad *articleDomain) Create(article *model.Article) *wrapper.Property {
	article.ID = uuid.New().String()
	result := ad.mgoRepo.InsertOne(article)
	return result
}
