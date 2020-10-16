package users

import (
	"github.com/JohanVong/GoAPI_beginner/datasources/mysql/usersdb"
)

const (
	queryInsertUser = "INSERT INTO users(firstname, lastname, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser    = "SELECT id, firstname, lastname, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET firstname=?, lastname=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindUser   = "SELECT id, firstname, lastname, email, date_created, status FROM users WHERE status=?;"
)

//TableName имя таблицы
func (User) TableName() string {
	return "users"
}

// Get user by ID
func (user *User) Get() error {
	var err error

	statement, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return err
	}
	defer statement.Close()

	result := statement.QueryRow(user.ID)
	if err = result.Scan(&user); err != nil {
		return err
	}
	return nil
}

// Create user in the DB
func (user *User) Create() error {
	var err error

	statement, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		user.Firstname,
		user.Lastname,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

// Update user in the DB
func (user *User) Update() error {
	statement, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		user.Firstname,
		user.Lastname,
		user.Email,
		user.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete user in the DB
func (user *User) Delete() error {
	statement, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.ID); err != nil {
		return err
	}
	return nil
}

// Find users by status
func (user *Users) Find(status string) error {
	statement, err := usersdb.Client.Prepare(queryFindUser)
	if err != nil {
		return err
	}
	defer statement.Close()

	rows, err := statement.Query(status)
	if err != nil {
		return err
	}
	defer rows.Close()

	if err = rows.Scan(&user); err != nil {
		return err
	}

	return nil
}
