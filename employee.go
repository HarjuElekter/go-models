package model

type Employee struct {
	ID        uint   `gorm:"primaryKey"`
	No        string `gorm:"type:varchar;size:20"`
	FirstName string `gorm:"type:varchar;size:30"`
	LastName  string `gorm:"type:varchar;size:30"`
}
