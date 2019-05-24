package domain

import (
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/patricksangian/go-rest-mux/helpers/utils"
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/model"
)

// UserDomain contains user properties and use cases
type userDomain struct {
	mgoRepo user.MongoRepositrory
}

// NewUserDomain acts as constructor
func NewUserDomain(mgoRepo user.MongoRepositrory) user.Domain {
	return &userDomain{
		mgoRepo: mgoRepo,
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
	result := ud.mgoRepo.InsertOne(user)
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