package services

import (
	"github.com/JohanVong/GoAPI_beginner/domain/users"
	"github.com/JohanVong/GoAPI_beginner/utils/errors"
)

// GetUser goes to Get() in Data Access Object '_dao' file
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser fills our struct with data via Save()
// in Data Access Object '_dao' file
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil

	// return &user, &errors.RestErr{
	// 	Status: http.StatusInternalServerError,
	// }
}
