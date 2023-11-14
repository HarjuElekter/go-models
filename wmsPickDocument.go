package model

import (
	"github.com/dainiauskas/go-types"
)

type PickDocument struct {
	ID         uint        `gorm:"primaryKey"`
	No         string      `gorm:"size:20;index"`
	Transfer   string      `gorm:"size:20;index"`
	Production string      `gorm:"size:20;index"`
	Project    string      `gorm:"size:20;index"`
	StartDate  *types.Date `gorm:"type:date"`
	EndDate    *types.Date `gorm:"type:date"`
	DueDate    *types.Date `gorm:"type:date"`
	Lines      uint        ``
	Assigned   uint        ``
	Users      []PickUser
}

func (PickDocument) TableName() string {
	return "wms_pick_documents"
}
