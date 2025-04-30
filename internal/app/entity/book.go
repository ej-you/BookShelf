package entity

type Book struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	UserID      string `gorm:"not null"`
	Title       string `gorm:"not null;size:150"`
	GenreID     string `gorm:"null"`
	Author      string `gorm:"null;size:150"`
	Year        string `gorm:"null"`
	Description string `gorm:"null;type:TEXT"`
	IsRead      bool   `gorm:"not null;default:0"`

	User  User  `gorm:"foreignKey:UserID"`
	Genre Genre `gorm:"foreignKey:GenreID"`
}

func (Book) TableName() string {
	return "book"
}
