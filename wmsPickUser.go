package model

import "time"

type PickUser struct {
	Date           time.Time `gorm:"type:date;primaryKey"`
	UserName       string    `gorm:"size:50;primaryKey"`
	Scanned        uint      ``
	PickDocumentID uint      ``
}

func (PickUser) TableName() string {
	return "wms_pick_users"
}
