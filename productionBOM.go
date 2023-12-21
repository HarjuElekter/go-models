package model

import "time"

type ProductionBOMLog struct {
	ID          uint       `gorm:"primaryKey" json:"id" query:"id" param:"id" header:"id"`
	Action      uint8      `json:"action" query:"action" param:"action" header:"action"`
	Type        string     `json:"type" query:"type" param:"type" header:"type"`
	UserName    string     `json:"username" query:"username" param:"username" header:"username"`
	BOMNo       string     `gorm:"size:20" json:"bomNo"  query:"bomNo" param:"bomNo" header:"bomNo"`
	ItemNo      string     `gorm:"size:20" json:"itemNo" query:"itemNo" param:"itemNo" header:"itemNo"`
	Quantity    float64    `json:"quantity" query:"quantity" param:"quantity" header:"quantity"`
	SystemId    string     `gorm:"size:50;index" json:"systemId"  query:"systemId" param:"systemId" header:"systemId"`
	OperationAt *time.Time `gorm:"type:datetime" json:"operationAt"  query:"operationAt" param:"operationAt" header:"operationAt"`
}

func (ProductionBOMLog) TableName() string {
	return "production_bom_log"
}
