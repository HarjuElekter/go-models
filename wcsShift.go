package model

import "time"

// Pamainų sąrašas
type ProdWorkShift struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"size:40;unique"`
	WCSShiftTimes []ProdWorkShiftTime
}

// Pamainos grafikas
type ProdWorkShiftTime struct {
	ID                uint                    `gorm:"primaryKey"`
	WCSShiftID        uint                    ``
	StartTime         string                  `gorm:"size:5"`
	EndTime           string                  `gorm:"size:5"`
	ValidFrom         *time.Time              `gorm:"type:date"`
	ValidTo           *time.Time              `gorm:"type:date"`
	WCSShiftSchedules []ProdWorkShiftSchedule ``
}

// Darbuotojų darbo grafikas
type ProdWorkShiftSchedule struct {
	ID             uint      `gorm:"primaryKey"`
	WCSShiftTimeID uint      ``
	Date           time.Time `gorm:"type:date"`
	EmployeeNo     string    `gorm:"size:20"`
}

// Darbuotojų darbo registracija
type ProdWorkShiftRegister struct {
	ID             uint              `gorm:"PrimaryKey"`
	WCSShiftTimeID uint              ``
	WCSShiftTime   ProdWorkShiftTime ``
	WorkStart      time.Time         `gorm:"type:time"`
	WorkEnd        *time.Time        `gorm:"type:time"`
}
