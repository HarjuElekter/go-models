package model

type Employee struct {
	ID        uint   `gorm:"primaryKey"`
	No        string `gorm:"size:20"`
	FirstName string `gorm:"size:30"`
	LastName  string `gorm:"size:30"`
}
