package usecase

import (
	"time"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
	"BookShelf/internal/app/repo"
	"BookShelf/internal/pkg/auth"
	"BookShelf/internal/pkg/password"
)

var _ UserUsecase = (*userUsecase)(nil)

// UserUsecase implementation.
type userUsecase struct {
	userRepoDB      repo.UserRepoDB
	tokenSigningKey []byte
	tokenTTL        time.Duration
}

func NewUserUsecase(
	userRepoDB repo.UserRepoDB,
	tokenSigningKey []byte,
	tokenTTL time.Duration) UserUsecase {

	return &userUsecase{
		userRepoDB:      userRepoDB,
		tokenSigningKey: tokenSigningKey,
		tokenTTL:        tokenTTL,
	}
}

// Sign up new user.
// Login and Password (not hashed) fields must be presented.
// Fill given struct and return it with new auth token.
func (u *userUsecase) SignUp(user *entity.User) (*entity.UserWithToken, error) {
	var err error
	// hash password
	user.Password, err = password.Encode(user.Password)
	if err != nil {
		return nil, err
	}

	// create user
	if err := u.userRepoDB.Create(user); err != nil {
		return nil, err
	}
	// generate auth token
	authToken, err := auth.NewToken(u.tokenSigningKey, u.tokenTTL, user.ID)
	if err != nil {
		return nil, err
	}

	return &entity.UserWithToken{
		User:      user,
		AuthToken: authToken,
	}, nil
}

// Login existing user.
// Login and Password (not hashed) fields must be presented.
// Fill given struct and return it with new auth token.
func (u *userUsecase) Login(user *entity.User) (*entity.UserWithToken, error) {
	// password entered by user
	enteredPasswd := user.Password

	// get user from DB with email
	if err := u.userRepoDB.GetByLogin(user); err != nil {
		return nil, err
	}

	// check entered password is correct
	if !password.IsCorrect(enteredPasswd, user.Password) {
		return nil, errors.ErrInvalidPassword
	}

	// generate auth token
	authToken, err := auth.NewToken(u.tokenSigningKey, u.tokenTTL, user.ID)
	if err != nil {
		return nil, err
	}

	return &entity.UserWithToken{
		User:      user,
		AuthToken: authToken,
	}, nil
}
