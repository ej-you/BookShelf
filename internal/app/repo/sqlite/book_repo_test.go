package sqlite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"BookShelf/internal/app/entity"
)

var (
	_userID        = "1"
	_genreID       = "1"
	_createdBookID = "1"

	_sortField      = "author"
	_sortOrder      = "desc"
	_filterType     = "read"
	_filterGenres   = []string{"antiophy", "russian classic"}
	_filterYearFrom = 0
	_filterYearTo   = 2100
)

func bookPrettyPrint(t *testing.T, book *entity.Book) {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s | ", book.ID))
	builder.WriteString(fmt.Sprintf("User: %s | ", book.UserID))
	builder.WriteString(fmt.Sprintf("%s | ", book.Title))
	if book.GenreID != nil {
		builder.WriteString(fmt.Sprintf("Genre: %s | ", book.Genre.Name))
	}
	if book.Author != nil {
		builder.WriteString(fmt.Sprintf("Author: %s | ", *book.Author))
	}
	if book.Year != nil {
		builder.WriteString(fmt.Sprintf("Year: %d | ", *book.Year))
	}
	builder.WriteString(fmt.Sprintf("IsRead: %v", book.IsRead))
	t.Log(builder.String())
}

func TestBook_Create(t *testing.T) {
	t.Log("Create new book")

	newBook := entity.Book{
		UserID:  _userID,
		Title:   "new_book_" + _uid,
		GenreID: &_genreID,
	}

	err := NewBookRepoDB(_dbStorage).Create(&newBook)
	require.NoError(t, err)

	_createdGenreID = newBook.ID
	t.Logf("New book: %+v", newBook)
}

func TestBook_GetList(t *testing.T) {
	t.Log("Get book list with sort and filters")

	books := &entity.BookList{
		UserID:         _userID,
		SortField:      &_sortField,
		SortOrder:      &_sortOrder,
		FilterType:     &_filterType,
		FilterYearFrom: &_filterYearFrom,
		FilterYearTo:   &_filterYearTo,
		FilterGenres:   _filterGenres,
	}

	err := NewBookRepoDB(_dbStorage).GetList(books)
	require.NoError(t, err)

	t.Logf("Found results: %d", len(books.Books))
	t.Log("Filtered and sorted books:")
	for _, book := range books.Books {
		bookPrettyPrint(t, &book)
	}
}

func TestBook_Update(t *testing.T) {
	t.Log("Update book by ID")

	year := 1949
	author := "George Orwell"
	existingBook := entity.Book{
		ID:      _createdBookID,
		UserID:  _userID,
		Title:   "1984 " + _uid,
		GenreID: &_genreID,
		Year:    &year,
		Author:  &author,
	}

	err := NewBookRepoDB(_dbStorage).Update(&existingBook)
	require.NoError(t, err)

	t.Logf("Updated book: %+v", existingBook)
}

func TestBook_GetByID(t *testing.T) {
	t.Log("Get book by ID")

	existingBook := entity.Book{ID: _createdBookID}

	err := NewBookRepoDB(_dbStorage).GetByID(&existingBook)
	require.NoError(t, err)

	t.Logf("Book: %+v", existingBook)
}

func TestBook_Remove(t *testing.T) {
	t.Log("Remove book by ID")

	existingBook := entity.Book{ID: _createdBookID}

	err := NewBookRepoDB(_dbStorage).Remove(&existingBook)
	require.NoError(t, err)

	t.Logf("Book with ID %s was deleted successfully", existingBook.ID)
}
