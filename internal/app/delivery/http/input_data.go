package http

type AuthInput struct {
	Login           string `form:"login" validate:"required,max=50"`
	Password        string `form:"password" validate:"required,min=8,max=30"`
	PasswordConfirm string `form:"password-confirm" validate:"-"` // for sign up only
}

type CreateGenreInput struct {
	Name string `form:"genre" validate:"required,max=50"`
}

type RemoveGenreInput struct {
	ID string `params:"genreID" validate:"required,numeric"`
}
