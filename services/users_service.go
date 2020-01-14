package services

import (
	"github.com/mritunjaykumar/users-api/domain/users"
	"github.com/mritunjaykumar/users-api/util/errors"
)

var (
	// UsersService implements UsersServiceInterface
	UsersService UsersServiceInterface = &usersService{}
)

// UsersServiceInterface defines behavior of UsersService
type UsersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
}

type usersService struct{}

func (u *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.ValidateEmail(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}

	// Validate userId
	if err := user.ValidateUserID(); err != nil {
		return nil, err
	}

	// Check if user exists in data store
	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}
