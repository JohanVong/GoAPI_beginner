package users

import (
	"fmt"

	"github.com/JohanVong/GoAPI_beginner/datasources/mysql/usersdb"
	"github.com/JohanVong/GoAPI_beginner/utils/errors"
)

const (
	queryInsertUser        = "INSERT INTO users(username, firstname, lastname, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?, ?);"
	queryGetUser           = "SELECT id, username, firstname, lastname, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser        = "UPDATE users SET firstname=?, lastname=?, email=?, status=?, password=? WHERE id=?;"
	queryDeleteUser        = "DELETE FROM users WHERE id=?;"
	queryFindUsersByStatus = "SELECT id, username, firstname, lastname, email, date_created, status FROM users WHERE status=?;"
)

// GetByID to get user by ID
func (user *User) GetByID() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(user.ID)
	if err = result.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return errors.DatabaseError(err.Error())
	}

	return nil
}

// Insert user in DB method
func (user *User) Insert() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.DatabaseError(err.Error())
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
	)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	user.ID = userID

	return nil
}

// Update user in the DB
func (user *User) Update() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.DatabaseError(err.Error())
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
		return errors.DatabaseError(err.Error())
	}

	return nil
}

// Delete user in the DB
func (user *User) Delete() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}
	defer statement.Close()

	if _, err = statement.Exec(user.ID); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// FindByStatus finds users by status
func (user *User) FindByStatus(status string) ([]User, *errors.CustomError) {
	var err error

	statement, err := usersdb.Client.Prepare(queryFindUsersByStatus)
	if err != nil {
		return nil, errors.DatabaseError(err.Error())
	}
	defer statement.Close()

	rows, err := statement.Query(status)
	if err != nil {
		return nil, errors.DatabaseError(err.Error())
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.DatabaseError(err.Error())
		}
		result = append(result, user)
	}

	if len(result) == 0 {
		return nil, errors.DatabaseError(fmt.Sprintf("No users matching status: %s", status))
	}

	return result, nil
}
