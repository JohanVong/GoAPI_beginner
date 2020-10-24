package users

import (
	"fmt"

	"github.com/JohanVong/GoAPI_beginner/datasources/mysql/usersdb"
	"github.com/JohanVong/GoAPI_beginner/utils/errors"
)

const (
	queryInsertUser           = "INSERT INTO users_db.users(username, firstname, lastname, email, date_created, status, password, is_admin) VALUES(?, ?, ?, ?, ?, ?, ?, ?);"
	queryGetUserByID          = "SELECT id, username, firstname, lastname, email, date_created, status, password, is_admin FROM users_db.users WHERE id=?;"
	queryGetUserByCredentials = "SELECT id, status, is_admin FROM users_db.users WHERE email=? AND password=?;"
	queryUpdateUser           = "UPDATE users_db.users SET firstname=?, lastname=?, email=?, status=?, password=? WHERE id=?;"
	queryDeleteUser           = "DELETE FROM users_db.users WHERE id=?;"
	queryFindUsersByStatus    = "SELECT id, username, firstname, lastname, email, date_created, status, is_admin FROM users_db.users WHERE status=?;"
)

// GetByID to get user by ID
func (user *User) GetByID() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryGetUserByID)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(user.ID)
	if err = result.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status, &user.Password, &user.IsAdmin); err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}

	return nil
}

// GetByCredentials to get user by email and password
func (user *User) GetByCredentials() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryGetUserByCredentials)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(user.Email, user.Password)
	if err = result.Scan(&user.ID, &user.Status, &user.IsAdmin); err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}

	return nil
}

// Insert user in DB method
func (user *User) Insert() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	defer statement.Close()

	result, err := statement.Exec(
		user.Username,
		user.Firstname,
		user.Lastname,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password,
		user.IsAdmin,
	)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	user.ID = uint(userID)

	return nil
}

// Update user in the DB
func (user *User) Update() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	defer statement.Close()

	_, err = statement.Exec(
		user.Firstname,
		user.Lastname,
		user.Email,
		user.Status,
		user.Password,
		user.ID,
	)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}

	return nil
}

// Delete user in the DB
func (user *User) Delete() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	defer statement.Close()

	if _, err = statement.Exec(user.ID); err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	return nil
}

// FindByStatus finds users by status
func (user *User) FindByStatus(status string) ([]User, *errors.CustomError) {
	var err error

	statement, err := usersdb.Client.Prepare(queryFindUsersByStatus)
	if err != nil {
		return nil, errors.TextError("An error in database occurred", err.Error())
	}
	defer statement.Close()

	rows, err := statement.Query(status)
	if err != nil {
		return nil, errors.TextError("An error in database occurred", err.Error())
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status, &user.IsAdmin); err != nil {
			return nil, errors.TextError("An error in database occurred", err.Error())
		}
		result = append(result, user)
	}

	if len(result) == 0 {
		return nil, errors.TextError("An error in database occurred", fmt.Sprintf("No users matching status: %s", status))
	}

	return result, nil
}
