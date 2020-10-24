package app

import (
	"github.com/JohanVong/GoAPI_beginner/controllers/ping"
	"github.com/JohanVong/GoAPI_beginner/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/auth/signup", users.CreateUser)
	router.POST("/auth/login", users.Login)

	router.GET("/users", users.SelfGetUser)
	router.PUT("/users", users.SelfUpdateUser)
	router.DELETE("/users", users.SelfDeleteUser)

	router.GET("/admin/users/:user_id", users.GetUserByID)
	router.PUT("/admin/users/:user_id", users.UpdateUserByID)
	router.DELETE("/admin/users/:user_id", users.DeleteUserByID)
	router.GET("/admin/users", users.SearchUsersByStatus)
}
