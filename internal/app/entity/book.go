package entity

type Book struct {
	ID          string  `gorm:"primaryKey;autoIncrement;type:INTEGER"`
	UserID      string  `gorm:"not null"`
	Title       string  `gorm:"not null;size:150"`
	GenreID     *string `gorm:"null"`
	Author      *string `gorm:"null;size:150"`
	Year        *int    `gorm:"null"`
	Description *string `gorm:"null;type:TEXT"`
	IsRead      bool    `gorm:"not null;default:0"`

	User  User  `gorm:"foreignKey:UserID"`
	Genre Genre `gorm:"foreignKey:GenreID"`
}

func (Book) TableName() string {
	return "book"
}

type BookList struct {
	UserID         string
	SortField      *string  `query:"sortField" validate:"omitempty,oneof=title year author"`
	SortOrder      *string  `query:"sortOrder" validate:"omitempty,oneof=asc desc"`
	FilterType     *string  `query:"filterType" validate:"omitempty,oneof=all read want"`
	FilterGenres   []string `query:"filterGenres" validate:"omitempty"`
	FilterYearFrom *int     `query:"filterYearFrom" validate:"omitempty,min=0,max=2100"`
	FilterYearTo   *int     `query:"filterYearTo" validate:"omitempty,min=0,max=2100"`
	Books          []Book
}
