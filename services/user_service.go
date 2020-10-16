package services

import (
	"github.com/JohanVong/GoAPI_beginner/domain/users"
	"github.com/JohanVong/GoAPI_beginner/utils/crypting"
	"github.com/JohanVong/GoAPI_beginner/utils/dateutil"
)

var (
	// UsersService is a variable for the entire
	// user service functionality
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, error)
	CreateUser(users.User) (*users.User, error)
	UpdateUser(bool, users.User) (*users.User, error)
	DeleteUser(int64) error
	SearchUsers(string) (*users.Users, error)
}

// GetUser goes to Get() in Data Access Object '_dao' file
func (s *usersService) GetUser(userID int64) (*users.User, error) {
	var err error
	result := &users.User{ID: userID}

	if err = result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser fills our struct with data via Save()
// in Data Access Object '_dao' file
func (s *usersService) CreateUser(user users.User) (*users.User, error) {
	var err error

	if err = user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = dateutil.GetNowDBFormat()
	user.Password = crypting.GetMd5(user.Password)

	if err = user.Create(); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates user via Update() in '_dao' file
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, error) {
	current, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.Firstname != "" {
			current.Firstname = user.Firstname
		}
		if user.Lastname != "" {
			current.Lastname = user.Lastname
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.Firstname = user.Firstname
		current.Lastname = user.Lastname
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

// DeleteUser removes user via Delete() in '_dao' file
func (s *usersService) DeleteUser(userID int64) error {
	user := &users.User{ID: userID}
	return user.Delete()
}

// Search gives users by status via Find() in '_dao' file
func (s *usersService) SearchUsers(status string) (*users.Users, error) {
	result := &users.Users{}
	if err := result.Find(status); err != nil {
		return nil, err
	}
	return result, nil
}
