package model

import (
	"time"
)

type RolePrivilege struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	RoleID    uint `json:"role_id" gorm:"column:role_id"`
	FeatureID uint `json:"feature_id" gorm:"column:feature_id"`
	Create    int8 `json:"create" gorm:"column:create"`
	Read      int8 `json:"read" gorm:"column:read"`
	Update    int8 `json:"update" gorm:"column:update"`
	Delete    int8 `json:"delete" gorm:"column:delete"`
}

func (RolePrivilege) TableName() string {
	return "role_privileges"
}
