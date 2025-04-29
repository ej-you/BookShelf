// Package entity contains all app entities
package entity

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"not null;size:50"`
	Email    string `gorm:"null;size:100"`
	Password string `gorm:"not null;size:255"`

	Books []Book `gorm:"foreignKey:UserID"`
}
