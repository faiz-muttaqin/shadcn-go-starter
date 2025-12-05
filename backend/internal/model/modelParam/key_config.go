package modelParam

import "time"

type KeyConfig struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	KeyType   string    `gorm:"column:key_type;type:varchar(255);collate:utf8mb4_unicode_ci" json:"key_type" form:"key_type"`
	Value     string    `gorm:"column:value;type:text;collate:utf8mb4_unicode_ci" json:"value" form:"value"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (KeyConfig) TableName() string {
	return "key_config"
}
