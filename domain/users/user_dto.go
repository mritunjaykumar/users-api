package users

import "github.com/mritunjaykumar/users-api/util/errors"

import "strings"

// User struct defines the user domain object
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// ValidateEmail validates the user object
func (u *User) ValidateEmail() *errors.RestErr {
	if strings.TrimSpace(strings.ToLower(u.Email)) == "" {
		return errors.BadRequest("invalid email address")
	}

	return nil
}

// ValidateUserID validates userID
func (u *User) ValidateUserID() *errors.RestErr {
	if u.ID <= 0 {
		return errors.BadRequest("userID can't be zero or negative")
	}

	return nil
}
