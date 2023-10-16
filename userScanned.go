package model

import (
	"time"

	"github.com/HarjuElekter/go-utils"
	"github.com/dainiauskas/go-types"
)

type UserScanned struct {
	Date         types.Date `gorm:"type:date;primaryKey" json:"-"`
	UserName     string     `gorm:"size:50;primaryKey" json:"username"`
	Lines        uint       `gorm:"smallint unsigned" json:"lines"`
	Start        *time.Time `gorm:"" json:"start"`
	End          *time.Time `gorm:"" json:"end"`
	RealDuration float64    `gorm:"tinyint unsigned" json:"duration"`
}

func (UserScanned) TableName() string {
	return "wms_user_scanned"
}

func (u *UserScanned) NameTrim() string {
	u.UserName = utils.TrimUserNamePrefix(u.UserName)
	return u.UserName
}
