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
	mysqlUsersPassword = "1"
	mysqlUsersHost     = "127.0.0.1:3306"
	mysqlUsersSchema   = "users_db"

	queryCreateUsersTable  = "CREATE TABLE `users_db`.`users` (`id` BIGINT(11) NOT NULL AUTO_INCREMENT, `username` VARCHAR(45) NOT NULL, `firstname` VARCHAR(45) NULL, `lastname` VARCHAR(45) NULL, `email` VARCHAR(45) NOT NULL, `date_created` DATETIME NULL, `status` VARCHAR(45) NULL, `password` VARCHAR(32) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `username_UNIQUE` (`username` ASC), UNIQUE INDEX `email_UNIQUE` (`email` ASC));"
	queryCreateUsersSchema = "CREATE SCHEMA `users_db` ;"
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

// CreateUsersDatabase to create database for future users
func CreateUsersDatabase() {
	var err error

	schemaStatement, err := Client.Prepare(queryCreateUsersSchema)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer schemaStatement.Close()

	tableStatement, err := Client.Prepare(queryCreateUsersTable)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer tableStatement.Close()

	_, err = schemaStatement.Exec()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = tableStatement.Exec()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
