// Package error provides base errors for application
// to use them in any layer.
package errors

import (
	goerrors "errors"
	"net/http"
)

var (
	ErrValidateData    = goerrors.New("validate data")               // HTTP code 400
	ErrInvalidPassword = goerrors.New("invalid password")            // HTTP code 401
	ErrConfirmPassword = goerrors.New("confirm password is invalid") // HTTP code 401
	ErrForbidden       = goerrors.New("forbidden")                   // HTTP code 403
	ErrNotFound        = goerrors.New("record not found")            // HTTP code 404
	ErrAlreadyExists   = goerrors.New("record already exists")       // HTTP code 409
	ErrParseAuthToken  = goerrors.New("internal error")              // HTTP code 500
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
	case goerrors.Is(err, ErrForbidden):
		return http.StatusForbidden
	case goerrors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case goerrors.Is(err, ErrAlreadyExists):
		return http.StatusConflict
	case goerrors.Is(err, ErrParseAuthToken):
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
