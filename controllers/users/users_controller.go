package users

import (
	"net/http"
	"strconv"

	"github.com/JohanVong/GoAPI_beginner/domain/users"
	"github.com/JohanVong/GoAPI_beginner/services"
	"github.com/JohanVong/GoAPI_beginner/utils/errors"

	"github.com/gin-gonic/gin"
)

// CreateUser for database
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser is used to get info about some user
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.RequestError("User id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
