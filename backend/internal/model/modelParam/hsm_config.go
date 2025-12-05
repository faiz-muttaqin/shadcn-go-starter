package modelParam

import "time"

type HsmConfig struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	HsmIP     string    `gorm:"column:hsm_ip;type:varchar(255);collate:utf8mb4_unicode_ci" json:"hsm_ip" form:"hsm_ip"`
	HsmPort   string    `gorm:"column:hsm_port;type:varchar(255);collate:utf8mb4_unicode_ci" json:"hsm_port" form:"hsm_port"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (HsmConfig) TableName() string {
	return "hsm_config"
}
