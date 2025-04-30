// Package error provides base errors for application
// to use them in any layer.
package errors

import goerrors "errors"

var (
	ErrNotFound      = goerrors.New("record not found")
	ErrAlreadyExists = goerrors.New("record already exists")
)
