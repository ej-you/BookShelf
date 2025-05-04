package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"BookShelf/internal/app/constants"
	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
	"BookShelf/internal/app/usecase"
	"BookShelf/internal/pkg/validator"
)

// Handlers for genre subroutes.
type genreHandler struct {
	genreUC usecase.GenreUsecase
	valid   validator.Validator
}

func newGenreHandler(genreUC usecase.GenreUsecase, valid validator.Validator) *genreHandler {
	return &genreHandler{
		genreUC: genreUC,
		valid:   valid,
	}
}

func (g *genreHandler) listHTML(ctx *fiber.Ctx) error {
	allGenres, err := g.genreUC.GetAll()
	if err != nil {
		return err
	}
	return ctx.Render("genre", fiber.Map{"genreList": allGenres})
}

func (g *genreHandler) create(ctx *fiber.Ctx) error {
	genreInput := &CreateGenreInput{}
	genre := &entity.Genre{}

	// parse form data
	if err := ctx.BodyParser(genreInput); err != nil {
		return err
	}
	// validate parsed data
	if err := g.valid.Validate(genreInput); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	// create new genre
	genre.Name = genreInput.Name
	err := g.genreUC.Create(genre)
	if err != nil {
		return err
	}
	return ctx.Redirect(constants.GenrePath, http.StatusSeeOther)
}

func (g *genreHandler) remove(ctx *fiber.Ctx) error {
	genreInput := &RemoveGenreInput{}
	genre := &entity.Genre{}

	// parse path params
	if err := ctx.ParamsParser(genreInput); err != nil {
		return err
	}
	// validate parsed data
	if err := g.valid.Validate(genreInput); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	// remove genre
	genre.ID = genreInput.ID
	err := g.genreUC.Remove(genre)
	if err != nil {
		return err
	}
	return ctx.Redirect(constants.GenrePath, http.StatusSeeOther)
}
