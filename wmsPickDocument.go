package model

import (
	"time"

	"github.com/dainiauskas/go-types"
)

type PickDocument struct {
	ID         uint        `gorm:"primaryKey"`
	No         string      `gorm:"size:20;index"`
	Transfer   string      `gorm:"size:20;index"`
	Production string      `gorm:"size:20;index"`
	Project    string      `gorm:"size:20;index"`
	Start      time.Time   `gorm:"type:datetime"`
	End        *time.Time  `gorm:"type:datetime"`
	DueDate    types.Date  `gorm:"type:date"`
	Lines      uint        ``
	Assigned   uint        ``
	Users      []*PickUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (PickDocument) TableName() string {
	return "wms_pick_documents"
}
