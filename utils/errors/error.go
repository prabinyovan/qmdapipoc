package errors

import (
	"strings"
)

type CustomError struct {
	ErrorCode             string
	UserFriendlyErrorCode string
	Error                 string
}

func NewCustomError(errorCode, UserFriendlyErrorCode string, error ...string) *CustomError {
	var customError = &CustomError{
		ErrorCode:             errorCode,
		UserFriendlyErrorCode: UserFriendlyErrorCode,
	}

	if error != nil {
		error_message := strings.TrimSpace(error[0])
		customError.Error = error_message
	}

	return customError
}
