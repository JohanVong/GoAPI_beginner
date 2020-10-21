package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	// _ is a blank pkg
	_ "github.com/go-sql-driver/mysql"
)

const (
	// Never load a production version credentials!!!
	mysqlUsersUsername = "Your Username"
	mysqlUsersPassword = "Your Password"
	mysqlUsersHost     = "Your Host"
	mysqlUsersSchema   = "Your Schema"
)

// Client is a mySQL database
var (
	Client *sql.DB

	username = mysqlUsersUsername
	password = mysqlUsersPassword
	host     = mysqlUsersHost
	schema   = mysqlUsersSchema
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database configured successfully")
}
