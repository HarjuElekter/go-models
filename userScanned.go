package model

import (
	"sync"
	"time"

	"github.com/HarjuElekter/go-utils"
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

func (p *PicksStat) IncDoc(wg *sync.WaitGroup) {
	p.Documents++
	wg.Done()
}

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
