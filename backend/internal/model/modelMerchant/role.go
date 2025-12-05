package modelMerchant

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName  string `json:"role_name" gorm:"column:role_name"`
	ClassName string `json:"class_name" gorm:"column:class_name"`
	Icon      string `json:"icon" gorm:"column:icon"`
	CreatedBy uint   `json:"created_by" gorm:"column:created_by"`
}

func (Role) TableName() string {
	return "roles"
}
