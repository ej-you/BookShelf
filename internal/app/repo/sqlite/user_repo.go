// Package repodb contains implementations for repo interfaces
// that use a DB as a data source.
package sqlite

import (
	goerrors "github.com/pkg/errors"
	"gorm.io/gorm"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
	"BookShelf/internal/app/repo"
)

var _ repo.UserRepoDB = (*userRepoDB)(nil)

// UserRepoDB implementation.
type userRepoDB struct {
	dbStorage *gorm.DB
}

func NewUserRepoDB(dbStorage *gorm.DB) repo.UserRepoDB {
	return &userRepoDB{
		dbStorage: dbStorage,
	}
}

// Get user by login.
// Login field must be presented.
// Fill given struct pointer value.
func (u *userRepoDB) GetByLogin(user *entity.User) error {
	err := u.dbStorage.Where("login = ?", user.Login).First(user).Error

	if goerrors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.ErrNotFound
	}
	return goerrors.Wrap(err, "get user by login")
}

// Create new user.
// All required fields must be presented.
func (u *userRepoDB) Create(user *entity.User) error {
	err := u.dbStorage.Create(user).Error

	if goerrors.Is(err, gorm.ErrDuplicatedKey) {
		err = errors.ErrAlreadyExists
	}
	return goerrors.Wrap(err, "create user")
}
