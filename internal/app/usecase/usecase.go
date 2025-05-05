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

type GenreUsecase interface {
	Create(genre *entity.Genre) error
	Remove(genre *entity.Genre) error
	GetAll() (*entity.GenreList, error)
}

type BookUsecase interface {
	Create(book *entity.Book) error
	Remove(book *entity.Book) error
	Update(book *entity.Book) error
	GetByID(book *entity.Book) error
	GetList(bookList *entity.BookList) error
}
