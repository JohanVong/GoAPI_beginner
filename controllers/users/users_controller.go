package users

import (
	"net/http"
	"strconv"

	"github.com/JohanVong/GoAPI_beginner/domain/users"
	"github.com/JohanVong/GoAPI_beginner/services"
	"github.com/JohanVong/GoAPI_beginner/utils/errors"

	"github.com/gin-gonic/gin"
)

// SelectServices is needed to select
// services if there is many of them
func SelectServices() {}

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
	var err error
	var customError *errors.CustomError
	var result *users.User

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON body",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	result, customError = services.UsersService.CreateUser(user)
	if customError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error on user creation",
			"error":   customError.Error,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User creation successful",
		"error":   nil,
		"data":    result.Marshall(c.GetHeader("X-Private") == "true"),
	})
}

// GetUser is used to get info about user
func GetUser(c *gin.Context) {
	var err error
	var customError *errors.CustomError
	var userID int64
	var user *users.User

	userID, err = getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse user ID",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	user, customError = services.UsersService.GetUser(userID)
	if customError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to retrieve a user",
			"error":   customError.Error,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully retrieved",
		"error":   nil,
		"data":    user.Marshall(c.GetHeader("X-Private") == "true"),
	})
}

// UpdateUser is used to alter user
func UpdateUser(c *gin.Context) {
	var user users.User
	var err error
	var customError *errors.CustomError

	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse user ID",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON body",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	user.ID = userID

	result, customError := services.UsersService.UpdateUser(user)
	if customError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update user",
			"error":   customError.Error,
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"error":   nil,
		"data":    result.Marshall(c.GetHeader("X-Private") == "true"),
	})
}

// DeleteUser is used to remove user from database
func DeleteUser(c *gin.Context) {
	var err error
	var customError *errors.CustomError

	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse user ID",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}
	if customError = services.UsersService.DeleteUser(userID); customError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete user",
			"error":   customError.Error,
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully deleted",
		"error":   nil,
		"data":    nil,
	})
}

// SearchUsersByStatus is used to find users by Query param
func SearchUsersByStatus(c *gin.Context) {
	var customError *errors.CustomError
	var users users.Users
	status := c.Query("status")

	users, customError = services.UsersService.FindUsersByStatus(status)
	if customError != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Users are not found",
			"error":   customError.Error,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users were successfully found",
		"error":   nil,
		"data":    users.Marshall(c.GetHeader("X-Private") == "true"),
	})
}
