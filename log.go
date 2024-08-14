package model

type Log struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	OS           string `gorm:"size:50" json:"os"`
	OSVersion    string `gorm:"size:254" json:"os_version"`
	AppName      string `gorm:"size:50" json:"app_name"`
	ErrorMessage string `gorm:"" json:"error_message"`
	Stack        string `gorm:"" json:"stack"`
}
