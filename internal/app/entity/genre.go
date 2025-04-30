package entity

type Genre struct {
	ID   string `gorm:"primaryKey;autoIncrement;type:INTEGER"`
	Name string `gorm:"not null;size:50;unique"`
}

func (Genre) TableName() string {
	return "genre"
}
