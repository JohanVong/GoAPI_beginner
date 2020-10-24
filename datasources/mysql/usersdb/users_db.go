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
	// Uncomment the constants below and assign your connection data to begin a work

	// MySQLusername = "Your username"
	// MySQLpassword = "Your password"
	// MySQLhost     = "Your host"

	queryCreateUsersSchema = "CREATE SCHEMA IF NOT EXISTS `users_db`;"
	queryCreateUsersTable  = "CREATE TABLE IF NOT EXISTS `users_db`.`users` (`id` BIGINT(11) NOT NULL AUTO_INCREMENT, `username` VARCHAR(45) NOT NULL, `firstname` VARCHAR(45) NULL, `lastname` VARCHAR(45) NULL, `email` VARCHAR(45) NOT NULL, `date_created` DATETIME NULL, `status` VARCHAR(45) NULL, `password` VARCHAR(32) NOT NULL, `is_admin` TINYINT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `username_UNIQUE` (`username` ASC), UNIQUE INDEX `email_UNIQUE` (`email` ASC));"
	queryCreateTokensTable = "CREATE TABLE IF NOT EXISTS `users_db`.`tokens` (`id` BIGINT(11) NOT NULL AUTO_INCREMENT, `user_id` BIGINT(11) NOT NULL, `token` VARCHAR(1000) NULL, PRIMARY KEY (`id`), UNIQUE INDEX `user_id_UNIQUE` (`user_id` ASC), CONSTRAINT `id` FOREIGN KEY (`user_id`) REFERENCES `users_db`.`users` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION);"
)

// Client is a mySQL database
var (
	Client *sql.DB

	username = MySQLusername
	password = MySQLpassword
	host     = MySQLhost
)

func init() {
	var err error

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8",
		username,
		password,
		host,
	)

	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	_, err = Client.Exec(queryCreateUsersSchema)
	if err != nil {
		panic(err)
	}

	_, err = Client.Exec(queryCreateUsersTable)
	if err != nil {
		panic(err)
	}

	_, err = Client.Exec(queryCreateTokensTable)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database configured successfully")
}
