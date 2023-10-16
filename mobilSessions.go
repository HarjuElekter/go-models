package model

import "time"

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
