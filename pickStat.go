package model

import (
	"github.com/dainiauskas/go-types"
)

type PicksStat struct {
	Date          types.Date     `gorm:"type:date;primaryKey" json:"-"`
	Documents     uint           `gorm:"type:smallint unsigned" json:"picksAll"`
	DocAssigned   uint           `gorm:"type:smallint unsigned" json:"picksAssigned"`
	Lines         uint           `gorm:"type:smallint unsigned" json:"pickLines"`
	LinesAssigned uint           `gorm:"type:smallint unsigned" json:"pickScanned"`
	ByUser        []*UserScanned `gorm:"-" json:"byUser"`
}

func (PicksStat) TableName() string {
	return "wms_picks_stat"
}
