package model

type Contact struct {
	Id        string `gorm:"column:id;type:varchar(16);primaryKey"`
	FirstName string `gorm:"column:first_name;type:varchar(15)"`
	LastName  string `gorm:"column:last_name;type:varchar(24)"`
}
