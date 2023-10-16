package model

import (
	"time"

	"github.com/dainiauskas/go-log"
	"gorm.io/gorm"
)

type MobileSession struct {
	ID                 int64      `gorm:"primaryKey"`
	SessionID          string     `gorm:"uniqueIndex" json:"sessionID"`
	UserName           string     `gorm:"" json:"username"`
	LoginTime          time.Time  `json:"loginTime"`
	LogoutTime         *time.Time `json:"logoutTime"`
	BrowserType        string     `gorm:"size:250" json:"browserType"`
	HostName           string     `gorm:"size:50"`
	NumberOfActivities int64      `json:"noOfActivities"`
	NumberOfParameters int64      `json:"noOfParameters"`
	OperInsert         int64      `json:"operInsert"`
	OperDelete         int64      `json:"operDelete"`
	OperModify         int64      `json:"operModify"`
}

func (MobileSession) TableName() string {
	return "mobile_sessions"
}

func (m *MobileSession) BeforeCreate(tx *gorm.DB) (err error) {
	m.LoginTime = time.Now()

	return
}

// Create - create user login record
func (m *MobileSession) Create(db *gorm.DB) (err error) {
	err = db.Create(m).Error
	if err != nil {
		log.Error("MobileSession Create error: %s", err)
	}

	return
}

// UpdatreLogout - set logout_time to time now
func (m *MobileSession) UpdateLogout(db *gorm.DB) (err error) {
	err = db.Where("session_id = ?", m.SessionID).Model(m).Update("logout_time", time.Now()).Error
	if err != nil {
		log.Error("MobileSession UpdateLogout error: %s", err)
	}

	return
}
