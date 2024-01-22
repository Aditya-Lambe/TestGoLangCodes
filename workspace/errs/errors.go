package errs

import (
	"net/http"
	//"golang.org/x/text/message"
)

type AppError struct {
	Code    int    `json:",omitempty"` //In case when your are not setting a particular field this will be omited from the response
	Message string `json:"message"`
}

/*
	This will take a message and will return a pointer to app error

Inside this we can simply return the address of app error
*/
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
