package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
)

var _createdGenreID = "1"

func TestGenre_Create(t *testing.T) {
	t.Log("Create new genre")

	newGenre := entity.Genre{
		Name: "new_genre_" + _uid,
	}

	err := NewGenreRepoDB(_dbStorage).Create(&newGenre)
	require.NoError(t, err)

	_createdGenreID = newGenre.ID
	t.Logf("New genre: %+v", newGenre)
}

func TestGenre_CreateDuplicate(t *testing.T) {
	t.Log("Try to create genre duplicate")

	newGenre := entity.Genre{
		Name: "new_genre_" + _uid,
	}

	err := NewGenreRepoDB(_dbStorage).Create(&newGenre)
	require.ErrorIs(t, err, errors.ErrAlreadyExists)
}

func TestGenre_GetAll(t *testing.T) {
	t.Log("Get all genres")

	allGenres, err := NewGenreRepoDB(_dbStorage).GetAll()
	require.NoError(t, err)

	t.Logf("All genres: %+v", allGenres)
}

func TestGenre_Remove(t *testing.T) {
	t.Log("Remove genre by ID")

	existingGenre := entity.Genre{ID: _createdGenreID}

	err := NewGenreRepoDB(_dbStorage).Remove(&existingGenre)
	require.NoError(t, err)

	t.Logf("Genre with ID %s was deleted successfully", existingGenre.ID)
}

func TestGenre_RemoveUnexisting(t *testing.T) {
	t.Log("Remove unexisting genre by ID")

	unexistingGenre := entity.Genre{ID: "0"}

	err := NewGenreRepoDB(_dbStorage).Remove(&unexistingGenre)
	require.NoError(t, err)

	t.Logf("Genre with ID %s was deleted (if it existed)", unexistingGenre.ID)
}
