package users

import (
	"regexp"
	"strings"

	"github.com/JohanVong/GoAPI_beginner/utils/errors"
)

const (
	// StatusActive is default status for all created users
	StatusActive = "active"
)

// User model
type User struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
	IsAdmin     bool   `json:"is_admin"`
}

// Users for array of User
type Users []User

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validate user
func (user *User) Validate() *errors.CustomError {

	user.Username = strings.TrimSpace(user.Username)
	user.Firstname = strings.TrimSpace(user.Firstname)
	user.Lastname = strings.TrimSpace(user.Lastname)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Username == "" {
		return errors.UsernameValidationError()
	}
	if user.Email == "" || isEmailValid(user.Email) == false {
		return errors.UserEmailValidationError()
	}
	user.Password = strings.TrimSpace(user.Password)
	if len(user.Password) < 6 {
		return errors.UserPassValidationError()
	}
	return nil
}

// ValidateOnUpdate user
func (user *User) ValidateOnUpdate() *errors.CustomError {
	user.Firstname = strings.TrimSpace(user.Firstname)
	user.Lastname = strings.TrimSpace(user.Lastname)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Status = strings.TrimSpace(strings.ToLower(user.Status))
	user.Password = strings.TrimSpace(user.Password)

	if user.Email != "" && isEmailValid(user.Email) == false {
		return errors.UserEmailValidationError()
	}

	if user.Password != "" && len(user.Password) < 6 {
		return errors.UserPassValidationError()
	}

	if user.Status != "" && user.Status != "active" && user.Status != "inactive" {
		return errors.UserStatusValidationError()
	}

	return nil
}

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
