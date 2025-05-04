package sqlite

import (
	goerrors "github.com/pkg/errors"
	"gorm.io/gorm"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
	"BookShelf/internal/app/repo"
)

var _ repo.GenreRepoDB = (*genreRepoDB)(nil)

// GenreRepoDB implementation.
type genreRepoDB struct {
	dbStorage *gorm.DB
}

func NewGenreRepoDB(dbStorage *gorm.DB) repo.GenreRepoDB {
	return &genreRepoDB{
		dbStorage: dbStorage,
	}
}

// Create new user.
// All required fields must be presented.
func (u *genreRepoDB) Create(genre *entity.Genre) error {
	err := u.dbStorage.Create(genre).Error

	if goerrors.Is(err, gorm.ErrDuplicatedKey) {
		return errors.ErrAlreadyExists
	}
	return err // err or nil
}

// Delete genre by its ID.
// ID field must be presented.
func (u *genreRepoDB) Remove(genre *entity.Genre) error {
	return u.dbStorage.Delete(&entity.Genre{}, genre.ID).Error
}

// Get all genres.
// Return slice of Genre structs.
func (u *genreRepoDB) GetAll() (*entity.GenreList, error) {
	genres := &entity.GenreList{}
	err := u.dbStorage.Find(genres).Error
	return genres, err
}
