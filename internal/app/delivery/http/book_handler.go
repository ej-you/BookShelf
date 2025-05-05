package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"BookShelf/internal/app/constants"
	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
	"BookShelf/internal/app/usecase"
	"BookShelf/internal/pkg/auth"
	"BookShelf/internal/pkg/validator"
)

// Handlers for book subroutes.
type bookHandler struct {
	bookUC          usecase.BookUsecase
	genreUC         usecase.GenreUsecase
	tokenSigningKey []byte
	valid           validator.Validator
}

func newBookHandler(
	bookUC usecase.BookUsecase,
	genreUC usecase.GenreUsecase,
	tokenSigningKey []byte,
	valid validator.Validator) *bookHandler {

	return &bookHandler{
		bookUC:          bookUC,
		genreUC:         genreUC,
		tokenSigningKey: tokenSigningKey,
		valid:           valid,
	}
}

func (b *bookHandler) listHTML(ctx *fiber.Ctx) error {
	bookList := &entity.BookList{}

	// parse query params
	if err := ctx.QueryParser(bookList); err != nil {
		return err
	}
	// validate parsed data
	if err := b.valid.Validate(bookList); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	// parse user ID from auth token
	authToken, ok := ctx.Locals(constants.LocalsKeyAuthToken).(string)
	if !ok {
		return errors.ErrParseAuthToken
	}
	userID, err := auth.ParseUserIDFromToken(b.tokenSigningKey, authToken)
	if err != nil {
		return err
	}

	// get all genres for genre filter
	allGenres, err := b.genreUC.GetAll()
	if err != nil {
		return fmt.Errorf("book list: %w", err)
	}

	bookList.UserID = userID
	// get book list according to sort and filter settings
	if err := b.bookUC.GetList(bookList); err != nil {
		return err
	}

	return ctx.Render("library", fiber.Map{
		"genreList": allGenres,
		"bookList":  bookList.Books,
		"total":     len(bookList.Books),
	})
}

func (b *bookHandler) createHTML(ctx *fiber.Ctx) error {
	// get all genres for genre select
	allGenres, err := b.genreUC.GetAll()
	if err != nil {
		return fmt.Errorf("create book: %w", err)
	}

	return ctx.Render("book_create", fiber.Map{
		"genreList": allGenres,
	})
}

func (b *bookHandler) editHTML(ctx *fiber.Ctx) error {
	bookIDInput := &BookIDInput{}
	book := &entity.Book{}

	// parse path params
	if err := ctx.ParamsParser(bookIDInput); err != nil {
		return err
	}
	// validate parsed data
	if err := b.valid.Validate(bookIDInput); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	// parse user ID from auth token
	authToken, ok := ctx.Locals(constants.LocalsKeyAuthToken).(string)
	if !ok {
		return errors.ErrParseAuthToken
	}
	userID, err := auth.ParseUserIDFromToken(b.tokenSigningKey, authToken)
	if err != nil {
		return err
	}

	// get book by given ID
	book.ID = bookIDInput.ID
	if err := b.bookUC.GetByID(book); err != nil {
		return fmt.Errorf("edit book: %w", err)
	}
	// if user want to edit a book that is not his own
	if userID != book.UserID {
		return fmt.Errorf("edit book: %w", errors.ErrForbidden)
	}

	// get all genres for genre select
	allGenres, err := b.genreUC.GetAll()
	if err != nil {
		return fmt.Errorf("edit book: %w", err)
	}

	return ctx.Render("book_edit", fiber.Map{
		"genreList": allGenres,
		"book":      book,
	})
}

func (b *bookHandler) create(ctx *fiber.Ctx) error {
	bookInput := &BookInput{}
	book := &entity.Book{}

	// parse form data
	if err := ctx.BodyParser(bookInput); err != nil {
		return err
	}
	// validate parsed data
	if err := b.valid.Validate(bookInput); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	// parse user ID from auth token
	authToken, ok := ctx.Locals(constants.LocalsKeyAuthToken).(string)
	if !ok {
		return errors.ErrParseAuthToken
	}
	userID, err := auth.ParseUserIDFromToken(b.tokenSigningKey, authToken)
	if err != nil {
		return err
	}

	// set input data to Book model
	book.UserID = userID
	b.bookInputToBook(bookInput, book)
	// create book
	if err := b.bookUC.Create(book); err != nil {
		return err
	}
	// redirect to library page
	return ctx.Redirect(constants.LibraryPath, http.StatusSeeOther)
}

func (b *bookHandler) edit(ctx *fiber.Ctx) error {
	bookIDInput := &BookIDInput{}
	bookInput := &BookInput{}

	// parse path params
	if err := ctx.ParamsParser(bookIDInput); err != nil {
		return err
	}
	// validate parsed data
	if err := b.valid.Validate(bookIDInput); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}
	// parse form data
	if err := ctx.BodyParser(bookInput); err != nil {
		return err
	}
	// validate parsed data
	if err := b.valid.Validate(bookInput); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	// parse user ID from auth token
	authToken, ok := ctx.Locals(constants.LocalsKeyAuthToken).(string)
	if !ok {
		return errors.ErrParseAuthToken
	}
	userID, err := auth.ParseUserIDFromToken(b.tokenSigningKey, authToken)
	if err != nil {
		return err
	}

	// set input data to Book model
	book := &entity.Book{}
	book.UserID = userID
	book.ID = bookIDInput.ID
	b.bookInputToBook(bookInput, book)
	// create book
	if err := b.bookUC.Update(book); err != nil {
		return err
	}
	// redirect to library page
	return ctx.Redirect(constants.LibraryPath, http.StatusSeeOther)
}

func (b *bookHandler) remove(ctx *fiber.Ctx) error {
	bookIDInput := &BookIDInput{}
	book := &entity.Book{}

	// parse path params
	if err := ctx.ParamsParser(bookIDInput); err != nil {
		return err
	}
	// validate parsed data
	if err := b.valid.Validate(bookIDInput); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	// parse user ID from auth token
	authToken, ok := ctx.Locals(constants.LocalsKeyAuthToken).(string)
	if !ok {
		return errors.ErrParseAuthToken
	}
	userID, err := auth.ParseUserIDFromToken(b.tokenSigningKey, authToken)
	if err != nil {
		return err
	}

	// get book by given ID
	book.ID = bookIDInput.ID
	if err := b.bookUC.GetByID(book); err != nil {
		return fmt.Errorf("remove book: %w", err)
	}
	// if user want to remove a book that is not his own
	if userID != book.UserID {
		return fmt.Errorf("remove book: %w", errors.ErrForbidden)
	}

	// remove book
	if err := b.bookUC.Remove(book); err != nil {
		return err
	}
	// redirect to library page
	return ctx.Redirect(constants.LibraryPath, http.StatusSeeOther)
}

// Copy all field values from BookInput to Book.
// Fill given Book struct by pointer.
func (b *bookHandler) bookInputToBook(bookInput *BookInput, book *entity.Book) {
	book.Title = bookInput.Title
	if bookInput.GenreID != "" {
		book.GenreID = &bookInput.GenreID
	}
	if bookInput.Author != "" {
		book.Author = &bookInput.Author
	}
	if bookInput.Year != 0 {
		book.Year = &bookInput.Year
	}
	if bookInput.Description != "" {
		book.Description = &bookInput.Description
	}
	if bookInput.Type == "read" {
		book.IsRead = true
	}
}
