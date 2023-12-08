package model

import "time"

type ProductionBOMLog struct {
	ID          uint       `gorm:"primaryKey"`
	Action      uint8      `json:"action"`
	Type        string     `json:"type"`
	UserName    string     `json:"username"`
	BOMNo       string     `gorm:"size:20" json:"bomNo"`
	ItemNo      string     `gorm:"size:20" json:"item_no"`
	Quantity    float64    `json:"quantity"`
	SystemId    string     `gorm:"size:50;index" json:"systemId"`
	OperationAt *time.Time `gorm:"type:datetime" json:"operationAt"`
}

func (ProductionBOMLog) TableName() string {
	return "production_bom_log"
}
