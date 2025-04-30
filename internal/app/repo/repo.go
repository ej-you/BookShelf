// Package repo contains interfaces of repositories
// and its implementations for all entities like
// DB repositories or cache repositories etc.
package repo

import (
	"BookShelf/internal/app/entity"
)

type UserRepoDB interface {
	Create(user *entity.User) error
	GetByLogin(user *entity.User) error
}

// type GenreRepoDB interface {
// }

// type BookRepoDB interface {
// }
