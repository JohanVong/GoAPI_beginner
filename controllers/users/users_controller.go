package users

import (
	"net/http"
	"strconv"

	"github.com/JohanVong/GoAPI_beginner/domain/users"
	"github.com/JohanVong/GoAPI_beginner/services"
	"github.com/JohanVong/GoAPI_beginner/utils/errors"

	"github.com/gin-gonic/gin"
)

func getUserID(userIDParam string) (int64, error) {
	var err error
	var userID int64

	userID, err = strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// CreateUser user in database
func CreateUser(c *gin.Context) {
	var user users.User
	var token string
	var err error
	var customError *errors.CustomError
	var result *users.User

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.JSONbindingError(err.Error()))
		return
	}

	result, customError = services.UsersService.CreateUser(user)
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.TextError("Error on user creation", customError.Error))
		return
	}

	token, customError = services.TokensService.CreateToken(uint(result.ID))
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.TextError("User is created without token", customError.Error))
		return
	}

	c.JSON(http.StatusCreated, errors.NoError(token))
}

// Login is used to get token with user email and pass
func Login(c *gin.Context) {
	var err error
	var token string
	var customError *errors.CustomError
	var user *users.User

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.JSONbindingError(err.Error()))
		return
	}

	user, customError = services.UsersService.GetUserByMailAndPass(user.Email, user.Password)
	if customError != nil {
		c.JSON(http.StatusNotFound, errors.TextError("Invalid credentials given", customError.Error))
		return
	}

	token, customError = services.TokensService.UpdateToken(user.ID)
	if customError != nil {
		c.JSON(http.StatusNotFound, errors.TextError("Failed to get a new token", customError.Error))
		return
	}

	c.JSON(http.StatusOK, errors.NoError(token))
}

// SelfGetUser is used to get info about user
func SelfGetUser(c *gin.Context) {
	var tokenstring string
	var customError *errors.CustomError
	var user *users.User

	tokenstring = c.GetHeader("X-Token")

	user, customError = services.TokensService.GetUserByToken(tokenstring)
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.GetUserIndexByTokenError(customError.Error))
		return
	}

	c.JSON(http.StatusOK, errors.NoError(user.Marshall(c.GetHeader("X-Private") == "true")))
}

// SelfUpdateUser is used to alter user
func SelfUpdateUser(c *gin.Context) {
	var tokenstring string
	var user users.User
	var operator *users.User
	var err error
	var customError *errors.CustomError

	tokenstring = c.GetHeader("X-Token")

	operator, customError = services.TokensService.GetUserByToken(tokenstring)
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.GetUserIndexByTokenError(customError.Error))
		return
	}

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.JSONbindingError(err.Error()))
		return
	}

	user.ID = operator.ID

	result, customError := services.UsersService.UpdateUser(user)
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.TextError("Failed to update user", customError.Error))
		return
	}

	c.JSON(http.StatusOK, errors.NoError(result.Marshall(c.GetHeader("X-Private") == "true")))
}

// SelfDeleteUser is used to remove user from database
func SelfDeleteUser(c *gin.Context) {
	var tokenstring string
	var customError *errors.CustomError
	var operator *users.User

	tokenstring = c.GetHeader("X-Token")

	operator, customError = services.TokensService.GetUserByToken(tokenstring)
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.GetUserIndexByTokenError(customError.Error))
		return
	}

	if customError = services.UsersService.DeleteUser(operator.ID); customError != nil {
		c.JSON(http.StatusBadRequest, errors.TextError("Failed to delete user", customError.Error))
		return
	}

	c.JSON(http.StatusOK, errors.NoError("null"))
}

// FUNCTIONS PROVIDED BELOW ARE FOR ADMIN !!!

// SearchUsersByStatus is used to find users by Query param
func SearchUsersByStatus(c *gin.Context) {
	var customError *errors.CustomError
	var operator *users.User
	var userArray users.Users
	var tokenstring string

	tokenstring = c.GetHeader("X-Token")

	operator, customError = services.TokensService.GetUserByToken(tokenstring)
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.GetUserIndexByTokenError(customError.Error))
		return
	}

	if operator.IsAdmin == false {
		c.JSON(http.StatusBadRequest, errors.AdminRightsException())
		return
	}

	status := c.Query("status")

	userArray, customError = services.UsersService.FindUsersByStatus(status)
	if customError != nil {
		c.JSON(http.StatusNotFound, errors.UserNotFoundError(customError.Error))
		return
	}

	c.JSON(http.StatusOK, errors.NoError(userArray.Marshall(c.GetHeader("X-Private") == "true")))
}

// GetUserByID is used to get info about user
func GetUserByID(c *gin.Context) {
	var tokenstring string
	var err error
	var customError *errors.CustomError
	var userID int64
	var user *users.User
	var operator *users.User

	tokenstring = c.GetHeader("X-Token")

	operator, customError = services.TokensService.GetUserByToken(tokenstring)
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.GetUserIndexByTokenError(customError.Error))
		return
	}

	if operator.IsAdmin == false {
		c.JSON(http.StatusBadRequest, errors.AdminRightsException())
		return
	}

	userID, err = getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.UserIndexParseError(customError.Error))
		return
	}

	user, customError = services.UsersService.GetUserByID(uint(userID))
	if customError != nil {
		c.JSON(http.StatusNotFound, errors.UserNotFoundError(customError.Error))
		return
	}

	c.JSON(http.StatusOK, errors.NoError(user.Marshall(c.GetHeader("X-Private") == "true")))
}

// UpdateUserByID is used to alter user
func UpdateUserByID(c *gin.Context) {
	var user users.User
	var err error
	var customError *errors.CustomError
	var tokenstring string
	var operator *users.User

	tokenstring = c.GetHeader("X-Token")

	operator, customError = services.TokensService.GetUserByToken(tokenstring)
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.GetUserIndexByTokenError(customError.Error))
		return
	}

	if operator.IsAdmin == false {
		c.JSON(http.StatusBadRequest, errors.AdminRightsException())
		return
	}

	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.UserIndexParseError(customError.Error))
		return
	}

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.JSONbindingError(err.Error()))
		return
	}

	user.ID = uint(userID)

	result, customError := services.UsersService.UpdateUser(user)
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.TextError("Failed to update user", customError.Error))
		return
	}

	c.JSON(http.StatusOK, errors.NoError(result.Marshall(c.GetHeader("X-Private") == "true")))
}

// DeleteUserByID is used to remove user from database
func DeleteUserByID(c *gin.Context) {
	var err error
	var customError *errors.CustomError
	var tokenstring string
	var operator *users.User

	tokenstring = c.GetHeader("X-Token")

	operator, customError = services.TokensService.GetUserByToken(tokenstring)
	if customError != nil {
		c.JSON(http.StatusBadRequest, errors.GetUserIndexByTokenError(customError.Error))
		return
	}

	if operator.IsAdmin == false {
		c.JSON(http.StatusBadRequest, errors.AdminRightsException())
		return
	}

	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.UserIndexParseError(customError.Error))
		return
	}

	if customError = services.UsersService.DeleteUser(uint(userID)); customError != nil {
		c.JSON(http.StatusBadRequest, errors.TextError("Failed to delete user", customError.Error))
		return
	}

	c.JSON(http.StatusOK, errors.NoError("null"))
}
