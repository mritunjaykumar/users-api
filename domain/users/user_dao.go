package users

import (
	"fmt"
	"github.com/mritunjaykumar/users-api/util/errors"
)

var (
	userDB = make(map[int64]*User) // mock datastore
)

// Get gets user from datastore
func (u *User) Get() *errors.RestErr {
	user := userDB[u.ID]

	if user == nil {
		return errors.NotFoundError(fmt.Sprintf("user with userID %d is not found", u.ID))
	}

	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.DateCreated = user.DateCreated

	return nil
}

// Save saves user into datastore
func (u *User) Save() *errors.RestErr {
	if userDB[u.ID] != nil {
		return errors.BadRequest("User already exist")
	}

	userDB[u.ID] = u

	return nil
}
