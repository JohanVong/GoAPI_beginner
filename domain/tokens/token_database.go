package tokens

import (
	"github.com/JohanVong/GoAPI_beginner/datasources/mysql/usersdb"
	"github.com/JohanVong/GoAPI_beginner/utils/errors"
)

const (
	queryGetUserIDbyToken = "SELECT id, user_id, token FROM users_db.tokens WHERE token=?;"
	queryInsertToken      = "INSERT INTO users_db.tokens(user_id, token) VALUES(?, ?);"
	queryUpdateToken      = "UPDATE users_db.tokens SET token=? WHERE user_id=?;"
	queryDeleteToken      = "DELETE FROM users_db.tokens WHERE user_id=?;"
)

// Insert token in DB method
func (token *Token) Insert() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryInsertToken)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	defer statement.Close()

	result, err := statement.Exec(
		token.UserID,
		token.Token,
	)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}

	tokenID, err := result.LastInsertId()
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	token.ID = uint(tokenID)

	return nil
}

// Retrieve to get userID by token
func (token *Token) Retrieve() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryGetUserIDbyToken)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(token.Token)
	if err = result.Scan(&token.ID, &token.UserID, &token.Token); err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}

	return nil
}

// Update user in the DB
func (token *Token) Update() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryUpdateToken)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	defer statement.Close()

	if _, err = statement.Exec(token.Token, token.UserID); err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}

	return nil
}

// Delete user in the DB
func (token *Token) Delete() *errors.CustomError {
	var err error

	statement, err := usersdb.Client.Prepare(queryDeleteToken)
	if err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	defer statement.Close()

	if _, err = statement.Exec(token.UserID); err != nil {
		return errors.TextError("An error in database occurred", err.Error())
	}
	return nil
}
