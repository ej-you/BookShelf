// Package error provides base errors for application
// to use them in any layer.
package errors

import (
	goerrors "errors"
	"net/http"
)

var (
	ErrValidateData    = goerrors.New("validate data")               // 400
	ErrInvalidPassword = goerrors.New("invalid password")            // 401
	ErrConfirmPassword = goerrors.New("confirm password is invalid") // 401
	ErrNotFound        = goerrors.New("record not found")            // 404
	ErrAlreadyExists   = goerrors.New("record already exists")       // 409
)

// Returns HTTP code for given error if given error is one of the error declared above.
func CodeByError(err error) int {
	switch {
	case goerrors.Is(err, ErrValidateData):
		return http.StatusBadRequest
	case goerrors.Is(err, ErrInvalidPassword):
		return http.StatusUnauthorized
	case goerrors.Is(err, ErrConfirmPassword):
		return http.StatusUnauthorized
	case goerrors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case goerrors.Is(err, ErrAlreadyExists):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
