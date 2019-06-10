package domain

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sangianpatrick/go-rest-mux/helpers/eventemitter"
	"github.com/sangianpatrick/go-rest-mux/helpers/utils"
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	articleModel "github.com/sangianpatrick/go-rest-mux/src/modules/article/model"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/model"
)

// userDomain contains user properties and use cases
type userDomain struct {
	mgoRepo user.MongoRepositrory
	emitter eventemitter.Emitter
}

// NewUserDomain acts as constructor
func NewUserDomain(mgoRepo user.MongoRepositrory, emitter eventemitter.Emitter) user.Domain {
	return &userDomain{
		mgoRepo: mgoRepo,
		emitter: emitter,
	}
}

// Create will create a new user.
func (ud *userDomain) Create(user *model.User) *wrapper.Property {
	encryptedPassword, err := utils.Encrypt([]byte(utils.SecretKey), user.Password)
	if err != nil {
		return wrapper.Error(http.StatusInternalServerError, err.Error())
	}
	user.ID = uuid.New().String()
	user.Password = encryptedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	result := ud.mgoRepo.InsertOne(user)
	if result.Success {
		ud.emitter.EmitEmailSender(
			"go-rest-mux",
			os.Getenv("EMAIL_USERNAME"),
			"[go-rest-mux] User Registration",
			fmt.Sprintf("Hai %s, you are registered", user.Name),
			[]string{user.Email},
		)
	}
	return result
}

// GetByID return user by spesific ID
func (ud *userDomain) GetByID(ID string) *wrapper.Property {
	result := ud.mgoRepo.FindByID(ID)
	return result
}

// GetByEmail return user profile with its password
func (ud *userDomain) GetByEmail(email string) *wrapper.Property {
	result := ud.mgoRepo.FindByEmail(email)
	return result
}

// GetAll returns list of user
func (ud *userDomain) GetAll(page int, size int) *wrapper.Property {
	result := ud.mgoRepo.FindAll(page, size)
	return result
}

func (ud *userDomain) Update(ID string, data *model.User) *wrapper.Property {
	data.UpdatedAt = time.Now()
	result := ud.mgoRepo.UpdateOne(ID, data)
	return result
}

func (ud *userDomain) Delete(ID string) *wrapper.Property {
	result := ud.mgoRepo.DeleteOne(ID)
	return result
}

func (ud *userDomain) CreateArticle(userID string, article *articleModel.Article) *wrapper.Property {
	retrievingUser := ud.GetByID(userID)
	if !retrievingUser.Success {
		return retrievingUser
	}
	user := retrievingUser.Data.(model.User)
	article.CreatedBy = user
	ud.emitter.EmitCreateArticle(article)
	result := wrapper.Data(http.StatusOK, nil, "user's article is on creating process")
	return result
}
