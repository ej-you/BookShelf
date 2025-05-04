// Package sqlite contains implementations for repo interfaces
// that use a SQLite DB as a data source.
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

// Create new user.
// All required fields must be presented.
func (u *userRepoDB) Create(user *entity.User) error {
	err := u.dbStorage.Create(user).Error

	if goerrors.Is(err, gorm.ErrDuplicatedKey) {
		return errors.ErrAlreadyExists
	}
	return err // err or nil
}

// Get user by login.
// Login field must be presented.
// Fill given struct pointer value.
func (u *userRepoDB) GetByLogin(user *entity.User) error {
	err := u.dbStorage.Where("login = ?", user.Login).First(user).Error

	if goerrors.Is(err, gorm.ErrRecordNotFound) {
		return errors.ErrNotFound
	}
	return err // err or nil
}
