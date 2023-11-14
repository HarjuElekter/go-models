package model

import "time"

type PickDocument struct {
	ID         uint      `gorm:"primaryKey"`
	No         string    `gorm:"size:20;index"`
	Transfer   string    `gorm:"size:20;index"`
	Production string    `gorm:"size:20;index"`
	Project    string    `gorm:"size:20;index"`
	Startdate  time.Time `gorm:"type:date"`
	EndDate    time.Time `gorm:"type:date"`
	DueDate    time.Time `grom:"type:date"`
	Lines      uint      ``
	Assigned   uint      ``
	Users      []PickUser
}

func (PickDocument) TableName() string {
	return "wms_pick_documents"
}
