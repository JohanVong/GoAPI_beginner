package app

import (
	"github.com/JohanVong/GoAPI_beginner/controllers/ping"
	"github.com/JohanVong/GoAPI_beginner/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("/internal/users/search", users.SearchUsersByStatus)
}
