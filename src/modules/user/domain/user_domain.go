package domain

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user"
)

// UserDomain contains user properties and use cases
type UserDomain struct {
	mgoRepo user.MongoRepositrory
}

// NewUserDomain acts as constructor
func NewUserDomain(mgoRepo user.MongoRepositrory) user.Domain {
	return &UserDomain{
		mgoRepo: mgoRepo,
	}
}

// GetByID return user by spesific ID
func (ud *UserDomain) GetByID(ID string) *wrapper.Property {
	result := ud.mgoRepo.FindByID(ID)
	return result
}

// GetProfile return user profile with its password
func (ud *UserDomain) GetProfile(email string) *wrapper.Property {
	result := ud.mgoRepo.FindByEmail(email)
	return result
}

// GetAll returns list of user
func (ud *UserDomain) GetAll(limit int, skip int) *wrapper.Property {
	result := ud.mgoRepo.FindAll(limit, skip)
	return result
}
