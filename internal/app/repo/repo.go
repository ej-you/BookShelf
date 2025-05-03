// Package repo contains interfaces of repositories for all entities.
// Its implementations like DB repositories, mocks,
// etc. are in sub-packages with the same names.
package repo

import (
	"BookShelf/internal/app/entity"
)

type UserRepoDB interface {
	Create(user *entity.User) error
	GetByLogin(user *entity.User) error
}

type GenreRepoDB interface {
	Create(genre *entity.Genre) error
	Remove(genre *entity.Genre) error
	GetAll() (*entity.GenreList, error)
	// GetAll(genres *entity.GenreList) error
}

// type BookRepoDB interface {
// }
