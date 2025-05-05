// Package auth provides functions to obtain auth tokens and parse
// data from them (e.g. user ID).
package auth

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

// Generate new auth token for user.
func NewToken(secretKey []byte, expireDuration time.Duration, userID string) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().UTC().Add(expireDuration).Unix(),
	})

	tokenString, err := tokenStruct.SignedString(secretKey)
	return tokenString, errors.Wrap(err, "obtain token")
}

// Parse user ID from given auth token.
func ParseUserIDFromToken(secretKey []byte, authToken string) (string, error) {
	// parse token object from token string representation
	jwtToken, err := jwt.Parse(authToken, func(_ *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("parse auth token: %w", err)
	}
	// parse user ID from token
	userID, err := jwtToken.Claims.GetSubject()
	if err != nil {
		return "", fmt.Errorf("parse user id from auth token: %w", err)
	}
	return userID, nil
}
