package users

import (
	"net/http"
	"strconv"

	"github.com/JohanVong/GoAPI_beginner/domain/users"
	"github.com/JohanVong/GoAPI_beginner/services"

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

// Create item in database
func Create(c *gin.Context) {
	var user users.User
	var err error
	var result *users.User

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON body",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	result, err = services.UsersService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User creation failed",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User creation successful",
		"error":   nil,
		"data":    result,
	})
}

// Get is used to get info about item
func Get(c *gin.Context) {
	var err error
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

	user, err = services.UsersService.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to retrieve a user",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully retrieved",
		"error":   nil,
		"data":    user,
	})
}

// Update is used to alter item
// partially or completely
func Update(c *gin.Context) {
	var user users.User
	var err error

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

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update user",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"error":   nil,
		"data":    result,
	})
}

// Delete is used to remove items from database
func Delete(c *gin.Context) {
	var err error

	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse user ID",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}
	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete user",
			"error":   err.Error(),
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

// Search is used to find items by Query param
func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUsers(status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Users are not found",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Users were successfully found",
		"error":   nil,
		"data":    users,
	})
}
