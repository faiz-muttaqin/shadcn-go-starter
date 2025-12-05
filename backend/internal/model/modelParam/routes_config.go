package modelParam

import "time"

type RoutesConfig struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Endpoint  string    `gorm:"column:endpoint;type:varchar(255);collate:utf8mb4_unicode_ci;not null" json:"endpoint" form:"endpoint"`
	URL       string    `gorm:"column:url;type:varchar(255);collate:utf8mb4_unicode_ci" json:"url" form:"url"`
	Status    int       `gorm:"column:status;default:1" json:"status" form:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (RoutesConfig) TableName() string {
	return "route_config"
}
