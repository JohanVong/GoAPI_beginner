package users

import (
	"strings"

	"github.com/JohanVong/GoAPI_beginner/utils/errors"
)

// User model
type User struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate user by email
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.RequestError("Invalid email address")
	}
	return nil
}
