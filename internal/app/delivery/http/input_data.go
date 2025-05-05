package http

type AuthInput struct {
	Login           string `form:"login" validate:"required,max=50"`
	Password        string `form:"password" validate:"required,min=8,max=30"`
	PasswordConfirm string `form:"password-confirm" validate:"-"` // for sign up only
}

type CreateGenreInput struct {
	Name string `form:"genre" validate:"required,max=20"`
}

type RemoveGenreInput struct {
	ID string `params:"genreID" validate:"required,numeric"`
}

type BookIDInput struct {
	ID string `params:"genreID" validate:"required,numeric"`
}

type BookInput struct {
	Title       string `form:"title" validate:"required,max=150"`
	GenreID     string `form:"genre" validate:"omitempty"`
	Author      string `form:"author" validate:"omitempty,max=150"`
	Year        int    `form:"year" validate:"omitempty"`
	Description string `form:"description" validate:"omitempty"`
	Type        string `form:"type" validate:"required,oneof=want read"`
}
