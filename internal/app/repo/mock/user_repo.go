// Package mock contains implementations for repo interfaces.
package mock

import (
	"bytes"

	"github.com/stretchr/testify/mock"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/repo"
	"BookShelf/internal/pkg/password"
)

var _ repo.UserRepoDB = (*UserRepoDB)(nil)

// UserRepoDB implementation.
type UserRepoDB struct {
	mock.Mock
}

func NewUserRepoDB() *UserRepoDB {
	return &UserRepoDB{}
}

func (m *UserRepoDB) GetByLogin(user *entity.User) error {
	args := m.Called(user)

	if args.Error(0) != nil {
		return args.Error(0)
	}

	// for invalid password create other hashed password
	if bytes.Equal(user.Password, []byte("invalid password")) {
		user.Password = []byte("password")
	}

	// hash password
	var err error
	user.Password, err = password.Encode(user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserRepoDB) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Returns matching func that match User objects by login.
func MatchUserByLogin(matchLogin string) any {
	return mock.MatchedBy(func(u *entity.User) bool {
		return u.Login == matchLogin
	})
}
