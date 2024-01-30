package model

type Employee struct {
	ID         uint   `gorm:"primaryKey"`
	No         string `gorm:"size:20;unique"`
	FirstName  string `gorm:"size:30"`
	LastName   string `gorm:"size:30"`
	Department string `gorm:"size:50;index"`
}
