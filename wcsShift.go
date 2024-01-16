package model

import "time"

// Pamainų sąrašas
type ProdWorkShift struct {
	ID                 uint                `gorm:"primaryKey" json:"id"`
	Name               string              `gorm:"size:40;unique" json:"name"`
	ProdWorkShiftTimes []ProdWorkShiftTime `json:"times"`
}

// Pamainos grafikas
type ProdWorkShiftTime struct {
	ID                     uint                    `gorm:"primaryKey"`
	ProdWorkShiftID        uint                    ``
	StartTime              string                  `gorm:"size:5"`
	EndTime                string                  `gorm:"size:5"`
	ValidFrom              *time.Time              `gorm:"type:date"`
	ValidTo                *time.Time              `gorm:"type:date"`
	ProdWorkShiftSchedules []ProdWorkShiftSchedule ``
}

// Darbuotojų darbo grafikas
type ProdWorkShiftSchedule struct {
	ID                  uint      `gorm:"primaryKey"`
	ProdWorkShiftTimeID uint      ``
	Date                time.Time `gorm:"type:date"`
	EmployeeNo          string    `gorm:"size:20"`
}

// Darbuotojų darbo registracija
type ProdWorkShiftRegister struct {
	ID                  uint              `gorm:"PrimaryKey"`
	ProdWorkShiftTimeID uint              ``
	ProdWorkShiftTime   ProdWorkShiftTime ``
	WorkStart           time.Time         `gorm:"type:time"`
	WorkEnd             *time.Time        `gorm:"type:time"`
}
