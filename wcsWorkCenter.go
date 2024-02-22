package model

import "time"

type ProdWorkCenter struct {
	ID          string `gorm:"size:20;primaryKey" json:"no"`
	Name        string `gorm:"size:100" json:"name"`
	Group       string `gorm:"size:10" json:"group"`
	MeasureUnit string `gorm:"size:10" json:"measureUnit"`
	Barcode     string `gorm:"size:50;index" json:"barcode"`
}

type ProdWCEmplRegister struct {
	ID               uint       `gorm:"primaryKey" json:"no"`
	ProdWorkCenterID string     ``
	EmployeeNo       string     `gorm:"size:20" json:"employeeNo"`
	WorkPackageNo    string     `gorm:"size:20" json:"workPackageNo"`
	Start            time.Time  `gorm:"type:time" json:"startTime"`
	End              *time.Time `gorm:"type:time" json:"endTime"`
}
