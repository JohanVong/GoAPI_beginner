package errors

// CustomError struct
type CustomError struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

// NoError to return successful response
func NoError(data interface{}) *CustomError {
	return &CustomError{
		Message: "Request successful",
		Error:   "null",
		Data:    data,
	}
}

// TextError to return text of error and error interface
func TextError(msgTxt string, errTxt string) *CustomError {
	return &CustomError{
		Message: msgTxt,
		Error:   errTxt,
		Data:    "null",
	}
}

// AdminRightsException to inform operator that he is not a superuser
func AdminRightsException() *CustomError {
	return &CustomError{
		Message: "Admin rights are required for this operation",
		Error:   "null",
		Data:    "null",
	}
}

// UserNotFoundError to return user not found response
func UserNotFoundError(text string) *CustomError {
	return &CustomError{
		Message: "No users found",
		Error:   text,
		Data:    "null",
	}
}

// JSONbindingError to return text Error about json validation failure
func JSONbindingError(text string) *CustomError {
	return &CustomError{
		Message: "Invalid JSON body",
		Error:   text,
		Data:    "null",
	}
}

// UsernameValidationError to return text Error about user validation failure
func UsernameValidationError() *CustomError {
	return &CustomError{
		Message: "Incorrect username was given",
		Error:   "Incorrect username format",
		Data:    "null",
	}
}

// UserEmailValidationError to return text Error about user validation failure
func UserEmailValidationError() *CustomError {
	return &CustomError{
		Message: "Incorrect email was given",
		Error:   "Incorrect email format",
		Data:    "null",
	}
}

// UserPassValidationError to return text Error about user validation failure
func UserPassValidationError() *CustomError {
	return &CustomError{
		Message: "Incorrect password was given",
		Error:   "Incorrect length or password format. Minimal length: 6",
		Data:    "null",
	}
}

// UserStatusValidationError to return text Error about user validation failure
func UserStatusValidationError() *CustomError {
	return &CustomError{
		Message: "Incorrect status was given",
		Error:   "Incorrect status format for user",
		Data:    "null",
	}
}

// UserIndexParseError to return text Error about user ID parse failure
func UserIndexParseError(text string) *CustomError {
	return &CustomError{
		Message: "Failed to parse user ID",
		Error:   text,
		Data:    "null",
	}
}

// GetUserIndexByTokenError to return text Error about user ID retrieve failure
func GetUserIndexByTokenError(text string) *CustomError {
	return &CustomError{
		Message: "A token may be invalid, expired or it belongs to user who does not exist anymore",
		Error:   text,
		Data:    "null",
	}
}
