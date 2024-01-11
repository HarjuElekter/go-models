package model

type Employee struct {
	ID        uint   `gorm:"primaryKey"`
	No        string `gorm:"length:20"`
	FirstName string `gorm:"length:30"`
	LastName  string `gorm:"length:30"`
}
