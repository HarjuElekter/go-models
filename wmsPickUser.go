package model

import "github.com/dainiauskas/go-types"

type PickUser struct {
	Date           types.Date `gorm:"type:date;primaryKey"`
	UserName       string     `gorm:"size:50;primaryKey"`
	Scanned        uint       ``
	PickDocumentID uint       `gorm:"primaryKey"`
}

func (PickUser) TableName() string {
	return "wms_pick_users"
}
