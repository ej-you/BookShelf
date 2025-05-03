package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
)

func TestUser_Create(t *testing.T) {
	t.Log("Create new user")

	newUser := entity.User{
		Login:    "new_user_" + _uid,
		Password: []byte("123"),
	}

	err := NewUserRepoDB(_dbStorage).Create(&newUser)
	require.NoError(t, err)

	t.Logf("New user: %+v", newUser)
}

func TestUser_CreateDuplicate(t *testing.T) {
	t.Log("Try to create user duplicate")

	newUser := entity.User{
		Login:    "new_user_" + _uid,
		Password: []byte("123"),
	}

	err := NewUserRepoDB(_dbStorage).Create(&newUser)
	require.ErrorIs(t, err, errors.ErrAlreadyExists)
}

func TestUser_GetByLogin(t *testing.T) {
	t.Log("Get user by login")

	existingUser := entity.User{Login: "new_user_" + _uid}

	err := NewUserRepoDB(_dbStorage).GetByLogin(&existingUser)
	require.NoError(t, err)

	t.Logf("Existing user: %+v", existingUser)
}

func TestUser_GetByLoginUnexisting(t *testing.T) {
	t.Log("Try to get user by unexisting login")

	unexistingUser := entity.User{Login: "new_user"}

	err := NewUserRepoDB(_dbStorage).GetByLogin(&unexistingUser)
	require.ErrorIs(t, err, errors.ErrNotFound)
}
