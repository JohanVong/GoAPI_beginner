package errors

// CustomError struct
type CustomError struct {
	Error string `json:"Error"`
}

// UsernameValidationError to return text Error about user validation failure
func UsernameValidationError() *CustomError {
	return &CustomError{
		Error: "Incorrect username format",
	}
}

// UserEmailValidationError to return text Error about user validation failure
func UserEmailValidationError() *CustomError {
	return &CustomError{
		Error: "Incorrect email format",
	}
}

// UserPassValidationError to return text Error about user validation failure
func UserPassValidationError() *CustomError {
	return &CustomError{
		Error: "Incorrect length or password format. Minimal length: 6",
	}
}

// UserStatusValidationError to return text Error about user validation failure
func UserStatusValidationError() *CustomError {
	return &CustomError{
		Error: "Incorrect status format for user",
	}
}

// DatabaseError to return text Error about user validation failure
func DatabaseError(text string) *CustomError {
	return &CustomError{
		Error: text,
	}
}
