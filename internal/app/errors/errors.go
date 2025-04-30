// Package error provides base errors for application
// to use them in any layer.
package errors

import goerrors "errors"

var (
	ErrInvalidPassword = goerrors.New("invalid password")      // 401
	ErrNotFound        = goerrors.New("record not found")      // 404
	ErrAlreadyExists   = goerrors.New("record already exists") // 409
)
