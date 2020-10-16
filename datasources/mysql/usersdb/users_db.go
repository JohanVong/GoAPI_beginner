package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	// _ is a blank pkg
	_ "github.com/go-sql-driver/mysql"
)

// USE LOCAL ENV VARIABLES FOR A GITHUB!!!
const (
	mysqlUsersUsername = "root"
	mysqlUsersPassword = "root"
	mysqlUsersHost     = "127.0.0.1:3306"
	mysqlUsersSchema   = "users_db"
)

// Client is a mySQL database
var (
	Client *sql.DB

	// username = os.Getenv(mysqlUsersUsername)
	// password = os.Getenv(mysqlUsersPassword)
	// host     = os.Getenv(mysqlUsersHost)
	// schema   = os.Getenv(mysqlUsersSchema)
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
