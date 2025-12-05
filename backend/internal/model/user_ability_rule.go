package model

import (
	"time"
)

type UserAbilityRule struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	RoleID    uint      `gorm:"column:role_id;index" json:"role_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Subject   string    `gorm:"column:subject;size:50" json:"subject"`
	Read      bool      `gorm:"column:read" json:"read"`
	Update    bool      `gorm:"column:update" json:"update"`
	Create    bool      `gorm:"column:create" json:"create"`
	Delete    bool      `gorm:"column:delete" json:"delete"`
	Role      *UserRole `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"role,omitempty"`
}

func (UserAbilityRule) TableName() string {
	return "user_ability_rules"
}
