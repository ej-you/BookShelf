package usecase

import (
	"github.com/pkg/errors"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/repo"
)

var _ BookUsecase = (*bookUsecase)(nil)

// BookUsecase implementation.
type bookUsecase struct {
	bookRepoDB repo.BookRepoDB
}

func NewBookUsecase(bookRepoDB repo.BookRepoDB) BookUsecase {
	return &bookUsecase{
		bookRepoDB: bookRepoDB,
	}
}

// Create new book.
// UserID and Title fields must be presented.
func (b *bookUsecase) Create(book *entity.Book) error {
	err := b.bookRepoDB.Create(book)
	return errors.Wrap(err, "create book")
}

// Delete book by its ID.
// ID field must be presented.
func (b *bookUsecase) Remove(book *entity.Book) error {
	err := b.bookRepoDB.Remove(book)
	return errors.Wrap(err, "remove book")
}

// Update all book fields with given data by giving book ID.
// ID, UserID and Title fields must be presented.
func (b *bookUsecase) Update(book *entity.Book) error {
	err := b.bookRepoDB.Update(book)
	return errors.Wrap(err, "update book")
}

// Get book by given ID with genre preloading.
// ID field must be presented.
// Fill given struct pointer value.
func (b *bookUsecase) GetByID(book *entity.Book) error {
	err := b.bookRepoDB.GetByID(book)
	return errors.Wrap(err, "get book by id")
}

// Get books by given sort and filters with genre preloading.
// UserID field must be presented.
// Fill given struct pointer value (Books field).
func (b *bookUsecase) GetList(bookList *entity.BookList) error {
	err := b.bookRepoDB.GetList(bookList)
	return errors.Wrap(err, "get book list")
}
