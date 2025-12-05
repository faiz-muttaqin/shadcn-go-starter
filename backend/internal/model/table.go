package model

import (
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Name string `form:"name" json:"name" gorm:"column:name"`
	Attr string `form:"attr" json:"attr" gorm:"column:attr"`
}

func (Table) TableName() string {
	return "tables"
}
