package model

import "time"

type Email struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Subject  string    `gorm:"column:subject" json:"subject"`
	Sender   string    `gorm:"column:sender" json:"sender"`
	Body     string    `gorm:"column:body" json:"body"`
	Received time.Time `gorm:"column:received" json:"received"`
}
