// Package entity contains all app entities.
package entity

type User struct {
	ID       string `gorm:"primaryKey;autoIncrement;type:INTEGER"`
	Login    string `gorm:"not null;size:50"`
	Password []byte `gorm:"not null;type:BLOB"`

	Books []Book `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "user"
}

type UserWithToken struct {
	User      *User
	AuthToken string
}
