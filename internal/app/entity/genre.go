package entity

type Genre struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;size:50;unique"`
}
