package model

import "github.com/dainiauskas/go-types"

type PickUser struct {
	Date           types.Date `gorm:"type:date;primaryKey" json:"-"`
	UserName       string     `gorm:"size:50;primaryKey" json:"username"`
	Scanned        uint       `json:"lines"`
	PickDocumentID uint       `gorm:"primaryKey" json:"-"`
}

func (PickUser) TableName() string {
	return "wms_pick_users"
}
