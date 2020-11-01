package services

import (
	"github.com/JohanVong/GoAPI_beginner/domain/tokens"
	"github.com/JohanVong/GoAPI_beginner/domain/users"
	"github.com/JohanVong/GoAPI_beginner/utils/errors"
	"github.com/JohanVong/GoAPI_beginner/utils/webtoken"
)

var (
	// TokensService is a variable for the entire tokens service functionality
	TokensService tokensServiceInterface = &tokensService{}
)

type tokensService struct{}

type tokensServiceInterface interface {
	GetUserByToken(tokenstring string) (*users.User, *errors.CustomError)
	CreateToken(update bool, userID uint) (string, *errors.CustomError)
	DeleteToken(userID uint) *errors.CustomError
}

// GetUserIDbyToken returns user by ID
func (t *tokensService) GetUserByToken(tokenstring string) (*users.User, *errors.CustomError) {
	var customError *errors.CustomError
	var token *tokens.Token
	var user *users.User

	token = &tokens.Token{Token: tokenstring}

	if customError = token.ValidateToken(); customError != nil {
		return nil, customError
	}

	if customError = token.Retrieve(); customError != nil {
		return nil, customError
	}

	user = &users.User{ID: token.UserID}

	if customError = user.GetByID(); customError != nil {
		return nil, customError
	}

	return user, nil
}

//CreateToken is a func which takes ID of user and creates a token for this user
func (t *tokensService) CreateToken(update bool, userID uint) (string, *errors.CustomError) {
	var customError *errors.CustomError
	var token tokens.Token
	var err error
	var newToken string

	newToken, err = webtoken.JWTCreate(userID, "global", 60)
	if err != nil {
		return "", errors.TextError("An error on JSON web token creation occurred", err.Error())
	}

	token.UserID = userID
	token.Token = newToken

	if update == true {
		if customError = token.Update(); customError != nil {
			return "", customError
		}
	} else {
		if customError = token.Insert(); customError != nil {
			return "", customError
		}
	}

	return newToken, nil
}

// DeleteToken is a func which takes userID and deletes a token corresponding to it
func (t *tokensService) DeleteToken(userID uint) *errors.CustomError {
	var customError *errors.CustomError
	var token tokens.Token

	token.UserID = userID

	if customError = token.Delete(); customError != nil {
		return customError
	}

	return nil
}
