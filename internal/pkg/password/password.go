// Package password provides functions to manage password hashes.
package password

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Encode given password.
// Returns encoded password like a hash.
func Encode(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return hash, errors.Wrap(err, "encode password")
}

// Check the given password (not hashed) is equal to its hash from DB
// Returns true, if passwords are equal.
func IsCorrect(password, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, password) == nil
}
