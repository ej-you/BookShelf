package usecase

import (
	"github.com/pkg/errors"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/repo"
)

var _ GenreUsecase = (*genreUsecase)(nil)

// GenreUsecase implementation.
type genreUsecase struct {
	genreRepoDB repo.GenreRepoDB
}

func NewGenreUsecase(genreRepoDB repo.GenreRepoDB) GenreUsecase {
	return &genreUsecase{
		genreRepoDB: genreRepoDB,
	}
}

// Create new genre.
// All required fields must be presented.
func (g *genreUsecase) Create(genre *entity.Genre) error {
	err := g.genreRepoDB.Create(genre)
	return errors.Wrap(err, "create genre")
}

// Delete genre by its ID.
// ID field must be presented.
func (g *genreUsecase) Remove(genre *entity.Genre) error {
	err := g.genreRepoDB.Remove(genre)
	return errors.Wrap(err, "remove genre")
}

// Get all genres.
// Fill given struct pointer value.
func (g *genreUsecase) GetAll() (*entity.GenreList, error) {
	genreList, err := g.genreRepoDB.GetAll()
	return genreList, errors.Wrap(err, "get all genres")
}
