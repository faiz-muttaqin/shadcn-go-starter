package modelMerchant

import "time"

type TerminalApproved struct {
	ID                  int64     `json:"ID" form:"ID" gorm:"column:ID;primarykey"`
	MID                 int64     `json:"MID" form:"MID" gorm:"column:MID"`
	SECRET              string    `json:"SECRET" form:"SECRET" gorm:"column:SECRET;type:varchar(255)"`
	STATUS              string    `json:"STATUS" form:"STATUS" gorm:"column:STATUS;type:varchar(255)"`
	LIMIT_AMOUNT        string    `json:"LIMIT_AMOUNT" form:"LIMIT_AMOUNT" gorm:"column:LIMIT_AMOUNT;size:30"`
	SERIAL_NUMBER       string    `json:"SERIAL_NUMBER" form:"SERIAL_NUMBER" gorm:"column:SERIAL_NUMBER;type:varchar(255)"`
	IMEI                string    `json:"IMEI" form:"IMEI" gorm:"column:IMEI;type:varchar(255)"`
	CREATED_AT          time.Time `json:"CREATED_AT" form:"CREATED_AT" gorm:"column:CREATED_AT;autoCreateTime"`
	UPDATED_AT          time.Time `json:"UPDATED_AT" form:"UPDATED_AT" gorm:"column:UPDATED_AT;autoUpdateTime"`
	EMAIL               string    `json:"EMAIL" form:"EMAIL" gorm:"column:EMAIL;type:varchar(255)"`
	PASSWORD            string    `json:"-" form:"-" gorm:"column:PASSWORD;type:varchar(255)"`
	PASSWORD_TEMP       string    `json:"-" form:"-" gorm:"column:PASSWORD_TEMP;type:varchar(255)"`
	IS_EMAIL_SENDED     bool      `json:"-" form:"-" gorm:"column:IS_EMAIL_SENDED;default:false"`
	IS_PASSWORD_RENEWED bool      `json:"-" form:"-" gorm:"column:IS_PASSWORD_RENEWED;default:false"`
	LAST_LOGIN          int64     `json:"-" form:"-" gorm:"column:LAST_LOGIN"`
	SESSION             string    `json:"-" form:"-" gorm:"column:SESSION;type:varchar(255)"`
	SESSION_EXPIRED     int64     `json:"-" form:"-" gorm:"column:SESSION_EXPIRED"`
}

func (TerminalApproved) TableName() string {
	return "terminals_approved"
}
