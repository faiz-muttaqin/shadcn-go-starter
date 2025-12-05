package model

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	ProfileImage   string         `json:"profile_image" gorm:"column:profile_image"`
	Fullname       string         `json:"fullname" gorm:"column:fullname"`
	Username       string         `json:"username" gorm:"column:username"`
	Phone          string         `json:"phone" gorm:"column:phone"`
	Email          string         `json:"email" gorm:"column:email"`
	Password       string         `json:"password" gorm:"column:password"`
	Type           int            `json:"type" gorm:"column:type"`
	Role           int            `json:"role" gorm:"column:role"`
	Status         int            `json:"status" gorm:"column:status"`
	CreateBy       uint           `json:"created_by" gorm:"column:created_by"`
	UpdateBy       uint           `json:"updated_by" gorm:"column:updated_by"`
	LastLogin      int64          `json:"last_login" gorm:"column:last_login"`
	LoginDelay     int64          `json:"login_delay" gorm:"column:login_delay"`
	Session        string         `json:"session" gorm:"column:session"`
	SessionExpired int64          `json:"session_expired" gorm:"column:session_expired"`
	IP             string         `json:"ip" gorm:"column:ip"`
}

func (Admin) TableName() string {
	return "admins"
}
