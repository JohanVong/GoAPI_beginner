package users

import (
	"fmt"

	"github.com/JohanVong/GoAPI_beginner/utils/date_util"
	"github.com/JohanVong/GoAPI_beginner/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Get user by ID
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.Firstname = result.Firstname
	user.Lastname = result.Lastname
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

// Save user in the DB
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.RequestError(fmt.Sprintf("user with email '%s' already exists", user.Email))
		}
		return errors.RequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	user.DateCreated = date_util.GetNowString()
	usersDB[user.ID] = user
	return nil
}
