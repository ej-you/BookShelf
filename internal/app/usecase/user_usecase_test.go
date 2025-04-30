package usecase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
	"BookShelf/internal/app/repo/mock"
)

var (
	_tokenSigningKey = []byte("sample")
	_tokenTTL        = 30 * time.Minute

	_newLogin       = "new_user"
	_existingLogin  = "existing_user"
	_password       = []byte("123")
	_invalid_passwd = []byte("invalid password")
)

func TestSignUp(t *testing.T) {
	t.Log("Sign up user")

	// setup user usecase
	mockRepo := mock.NewUserRepoDB()
	userUC := NewUserUsecase(mockRepo, _tokenSigningKey, _tokenTTL)

	// create test users
	var newUser = &entity.User{ID: 1, Login: _newLogin, Password: _password}
	var existingUser = &entity.User{ID: 2, Login: _existingLogin, Password: _password}

	// setup mock expectations
	mockRepo.On("Create", mock.MatchUserByLogin(_newLogin)).Return(nil)
	mockRepo.On("Create", mock.MatchUserByLogin(_existingLogin)).Return(errors.ErrAlreadyExists)

	newUserWithToken, err := userUC.SignUp(newUser)
	require.NoError(t, err)
	t.Logf("New registred user: %+v", newUserWithToken.User)
	t.Logf("New registred user auth token: %s", newUserWithToken.AuthToken)

	_, err = userUC.SignUp(existingUser)
	require.ErrorIs(t, err, errors.ErrAlreadyExists)

	// check that all mock expectations was used
	mockRepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	t.Log("Login user")

	// setup user usecase
	mockRepo := mock.NewUserRepoDB()
	userUC := NewUserUsecase(mockRepo, _tokenSigningKey, _tokenTTL)

	// create test users
	var existingUser = &entity.User{ID: 1, Login: _existingLogin, Password: _password}
	var newUser = &entity.User{ID: 2, Login: _newLogin, Password: _password}
	var withInvalidPassword = &entity.User{ID: 3, Login: _existingLogin, Password: _invalid_passwd}

	// setup mock expectations
	mockRepo.On("GetByLogin", mock.MatchUserByLogin(_existingLogin)).Return(nil)
	mockRepo.On("GetByLogin", mock.MatchUserByLogin(_newLogin)).Return(errors.ErrNotFound)

	// without error
	existingUserWithToken, err := userUC.Login(existingUser)
	require.NoError(t, err)
	t.Logf("Logged in user: %+v", existingUserWithToken.User)
	t.Logf("Logged in user auth token: %s", existingUserWithToken.AuthToken)

	// unexisting user
	_, err = userUC.Login(newUser)
	require.ErrorIs(t, err, errors.ErrNotFound)
	// invalid password
	_, err = userUC.Login(withInvalidPassword)
	require.ErrorIs(t, err, errors.ErrInvalidPassword)

	// check that all mock expectations was used
	mockRepo.AssertExpectations(t)
}
