package model

type MobileUser struct {
	ID       int64  `gorm:"primaryKey"`
	Username string `gorm:"size:50;uniqueIndex"`
}
