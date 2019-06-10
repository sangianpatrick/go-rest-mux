package user

import (
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	articleModel "github.com/sangianpatrick/go-rest-mux/src/modules/article/model"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/model"
)

// Domain contains user's behavior.
type Domain interface {
	Create(user *model.User) *wrapper.Property
	GetByEmail(email string) *wrapper.Property
	GetByID(id string) *wrapper.Property
	GetAll(page int, size int) *wrapper.Property
	Update(ID string, data *model.User) *wrapper.Property
	Delete(ID string) *wrapper.Property
	CreateArticle(userID string, article *articleModel.Article) *wrapper.Property
}
