package app

import (
	"github.com/JohanVong/GoAPI_beginner/controllers/users"
)

func mapUrls() {
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", users.SearchUser)
}
