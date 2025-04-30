package repo

import (
	goerrors "github.com/pkg/errors"
	"gorm.io/gorm"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
)

var _ UserRepoDB = (*userRepoDB)(nil)

// UserRepoDB implementation.
type userRepoDB struct {
	dbStorage *gorm.DB
}

func NewUserRepoDB(dbStorage *gorm.DB) UserRepoDB {
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
// All required fields must be presented. All optional fields are optional.
func (u *userRepoDB) Create(user *entity.User) error {
	err := u.dbStorage.Create(user).Error

	if goerrors.Is(err, gorm.ErrDuplicatedKey) {
		err = errors.ErrAlreadyExists
	}
	return goerrors.Wrap(err, "create user")
}
