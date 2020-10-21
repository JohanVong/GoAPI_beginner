package services

import (
	"github.com/JohanVong/GoAPI_beginner/domain/users"
	"github.com/JohanVong/GoAPI_beginner/utils/crypting"
	"github.com/JohanVong/GoAPI_beginner/utils/dateutil"
	"github.com/JohanVong/GoAPI_beginner/utils/errors"
)

var (
	// UsersService is a variable for the entire
	// user service functionality
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.CustomError)
	CreateUser(users.User) (*users.User, *errors.CustomError)
	UpdateUser(users.User) (*users.User, *errors.CustomError)
	DeleteUser(int64) *errors.CustomError
	FindUsersByStatus(string) (users.Users, *errors.CustomError)
}

// GetUser goes to Get() in Data Access Object '_dao' file
func (s *usersService) GetUser(userID int64) (*users.User, *errors.CustomError) {
	var customError *errors.CustomError
	user := &users.User{ID: userID}

	if customError = user.GetByID(); customError != nil {
		return nil, customError
	}
	return user, nil
}

// CreateUser fills our struct with data via Save()
// in Data Access Object '_dao' file
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.CustomError) {
	var customError *errors.CustomError

	if customError = user.Validate(); customError != nil {
		return nil, customError
	}

	user.Status = users.StatusActive
	user.DateCreated = dateutil.GetNowDBFormat()
	user.Password = crypting.GetMd5(user.Password)

	if customError = user.Insert(); customError != nil {
		return nil, customError
	}

	return &user, nil
}

// UpdateUser updates user via Update() in '_dao' file
func (s *usersService) UpdateUser(user users.User) (*users.User, *errors.CustomError) {
	var customError *errors.CustomError

	if customError = user.ValidateOnUpdate(); customError != nil {
		return nil, customError
	}

	oldUser, customError := s.GetUser(user.ID)
	if customError != nil {
		return nil, customError
	}

	if user.Firstname != "" {
		oldUser.Firstname = user.Firstname
	}
	if user.Lastname != "" {
		oldUser.Lastname = user.Lastname
	}
	if user.Email != "" {
		oldUser.Email = user.Email
	}
	if user.Status != "" {
		oldUser.Status = user.Status
	}
	if user.Password != "" {
		oldUser.Password = crypting.GetMd5(user.Password)
	}

	if err := oldUser.Update(); err != nil {
		return nil, err
	}

	return oldUser, nil
}

// DeleteUser removes user via Delete() in '_dao' file
func (s *usersService) DeleteUser(userID int64) *errors.CustomError {
	var customError *errors.CustomError
	user := &users.User{ID: userID}

	if customError = user.Delete(); customError != nil {
		return customError
	}

	return nil
}

// Search gives users by status via Find() in '_dao' file
func (s *usersService) FindUsersByStatus(status string) (users.Users, *errors.CustomError) {

	databaseObject := &users.User{}

	return databaseObject.FindByStatus(status)
}
