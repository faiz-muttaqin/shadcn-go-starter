package modelParam

import (
	"gorm.io/gorm"
)

type SMTP struct {
	gorm.Model
	HOST     string `gorm:"column:host;type:varchar(255)" json:"host" form:"host"`
	PORT     int    `gorm:"column:port" json:"port" form:"port"`
	EMAIL    string `gorm:"column:email;type:varchar(255)" json:"email" form:"email"`
	PASSWORD string `gorm:"column:password;type:varchar(255)" json:"password" form:"password"`
	SENDER   string `gorm:"column:sender;type:varchar(255)" json:"sender" form:"sender"`
}

func (SMTP) TableName() string {
	return "smtp"
}
