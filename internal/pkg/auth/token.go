// Package auth provides functions to obtain auth tokens and parse
// data from them (e.g. user ID).
package auth

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

// Generate new auth token for user.
func NewToken(secretKey []byte, expireDuration time.Duration, userID int) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().UTC().Add(expireDuration).Unix(),
	})

	tokenString, err := tokenStruct.SignedString(secretKey)
	return tokenString, errors.Wrap(err, "obtain token")
}

// Parse user ID from given auth token.
func ParseUserIDFromToken(secretKey any, authToken string) (int, error) {
	// parse token object from token string representation
	jwtToken, err := jwt.Parse(authToken, func(_ *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return 0, fmt.Errorf("parse auth token: %w", err)
	}
	// parse user ID from token as string
	stringUserID, err := jwtToken.Claims.GetSubject()
	if err != nil {
		return 0, fmt.Errorf("parse user id from auth token: %w", err)
	}
	// convert string user ID to int
	intUserID, err := strconv.Atoi(stringUserID)
	if err != nil {
		return 0, fmt.Errorf("parse user id value: %w", err)
	}
	return intUserID, nil
}
