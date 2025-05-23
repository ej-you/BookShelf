package sqlite

import (
	"fmt"

	goerrors "github.com/pkg/errors"
	"gorm.io/gorm"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
	"BookShelf/internal/app/repo"
)

const (
	_filterYearFromMin = 0
	_filterYearToMax   = 2100
)

var _ repo.BookRepoDB = (*bookRepoDB)(nil)

// GenreRepoDB implementation.
type bookRepoDB struct {
	dbStorage *gorm.DB
}

func NewBookRepoDB(dbStorage *gorm.DB) repo.BookRepoDB {
	return &bookRepoDB{
		dbStorage: dbStorage,
	}
}

// Create new book.
// UserID and Title fields must be presented.
func (u *bookRepoDB) Create(book *entity.Book) error {
	err := u.dbStorage.Create(book).Error

	if goerrors.Is(err, gorm.ErrDuplicatedKey) {
		return errors.ErrAlreadyExists
	}
	return err // err or nil
}

// Delete book by its ID.
// ID field must be presented.
func (u *bookRepoDB) Remove(book *entity.Book) error {
	return u.dbStorage.Delete(&entity.Book{}, book.ID).Error
}

// Update all book fields with given data by giving book ID.
// ID, UserID and Title fields must be presented.
func (u *bookRepoDB) Update(book *entity.Book) error {
	return u.dbStorage.Save(book).Error
}

// Get book by given ID with genre preloading.
// ID field must be presented.
// Fill given struct pointer value.
func (u *bookRepoDB) GetByID(book *entity.Book) error {
	err := u.dbStorage.Preload("Genre").Where("id = ?", book.ID).First(book).Error

	if goerrors.Is(err, gorm.ErrRecordNotFound) {
		return errors.ErrNotFound
	}
	return err // err or nil
}

// Get books by given sort and filters with genre preloading.
// UserID field must be presented.
// Fill given struct pointer value (Books field).
func (u *bookRepoDB) GetList(bookList *entity.BookList) error {
	// base select params
	selectQuery := u.dbStorage.
		Table("book").
		Distinct("book.*").
		Joins("LEFT JOIN genre ON book.genre_id = genre.id").
		Where("user_id = ?", bookList.UserID)
	selectQuery = u.addSort(selectQuery, bookList)
	selectQuery = u.addFilter(selectQuery, bookList)
	// do select query
	return selectQuery.Preload("Genre").Find(&bookList.Books).Error
}

func (u *bookRepoDB) addSort(selectQuery *gorm.DB, bookList *entity.BookList) *gorm.DB {
	// apply sort by title if sort field is not specified
	if bookList.SortField == nil {
		titleField := "title"
		bookList.SortField = &titleField
	}
	// set asc order if sort order is not specified
	if bookList.SortOrder == nil {
		ascOrder := "asc"
		bookList.SortOrder = &ascOrder
	}
	return selectQuery.Order(fmt.Sprintf("%s %s", *bookList.SortField, *bookList.SortOrder))
}

func (u *bookRepoDB) addFilter(selectQuery *gorm.DB, bookList *entity.BookList) *gorm.DB {
	if bookList.FilterType != nil {
		switch *bookList.FilterType {
		case "read":
			selectQuery = selectQuery.Where("is_read = 1")
		case "want":
			selectQuery = selectQuery.Where("is_read = 0")
		}
	}
	if bookList.FilterYearFrom != nil && *bookList.FilterYearFrom != _filterYearFromMin {
		selectQuery = selectQuery.Where("year >= ?", *bookList.FilterYearFrom)
	}
	if bookList.FilterYearTo != nil && *bookList.FilterYearTo != _filterYearToMax {
		selectQuery = selectQuery.Where("year <= ?", *bookList.FilterYearTo)
	}
	// have at least one genre from given genres
	if len(bookList.FilterGenres) > 0 {
		selectQuery = selectQuery.Where(`genre.name IN ?`, bookList.FilterGenres)
	}
	return selectQuery
}
