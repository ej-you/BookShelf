// Package usecase contains interfaces of usecases
// and its implementations for all entities.
package usecase

import (
	"BookShelf/internal/app/entity"
)

type UserUsecase interface {
	SignUp(user *entity.User) (*entity.UserWithToken, error)
	Login(user *entity.User) (*entity.UserWithToken, error)
}
