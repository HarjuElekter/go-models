package model

import (
	"strings"
	"time"

	"github.com/dainiauskas/go-log"
	"gorm.io/gorm"
)

type MobileSession struct {
	ID          int64      `gorm:"primaryKey"`
	SessionID   string     `gorm:"uniqueIndex" json:"sessionID"`
	UserName    string     `gorm:"" json:"username"`
	LoginTime   time.Time  `json:"loginTime"`
	LogoutTime  *time.Time `json:"logoutTime"`
	BrowserType string     `gorm:"size:250" json:"browserType"`
	HostName    string     `gorm:"size:50"`
}

func (MobileSession) TableName() string {
	return "wms_mobile_sessions"
}

func (m *MobileSession) BeforeCreate(tx *gorm.DB) (err error) {
	m.LoginTime = time.Now()
	m.UserName = strings.ToUpper(m.UserName)

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
