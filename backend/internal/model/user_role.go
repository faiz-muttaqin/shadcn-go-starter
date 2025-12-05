package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	ID           uint              `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt    time.Time         `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time         `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt    `gorm:"index;column:deleted_at" json:"deleted_at"`
	Title        string            `json:"title" gorm:"column:title"`
	Name         string            `json:"name" gorm:"column:name"`
	Icon         string            `json:"icon" gorm:"column:icon"`
	CreatedBy    uint              `json:"created_by" gorm:"column:created_by"`
	AbilityRules []UserAbilityRule `gorm:"foreignKey:RoleID;references:ID" json:"ability_rules"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
