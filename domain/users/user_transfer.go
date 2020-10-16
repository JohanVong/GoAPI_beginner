package users

import (
	"strings"
)

const (
	// StatusActive is default status for all created users
	StatusActive = "active"
)

// User model
type User struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Users for array of User
type Users []User

// Validate user by email
func (user *User) Validate() error {
	var err error

	user.Firstname = strings.TrimSpace(user.Firstname)
	user.Lastname = strings.TrimSpace(user.Lastname)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return err
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return err
	}
	return nil
}
