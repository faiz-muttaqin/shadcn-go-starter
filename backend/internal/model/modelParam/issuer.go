package modelParam

import "time"

type Issuer struct {
	ID               int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IssuerName       string    `gorm:"column:issuer_name;type:varchar(255);collate:utf8mb4_unicode_ci" json:"issuer_name" form:"issuer_name"`
	IssuerType       string    `gorm:"column:issuer_type;type:varchar(255);collate:utf8mb4_unicode_ci" json:"issuer_type" form:"issuer_type"`
	IssuerURLService string    `gorm:"column:issuer_url_service;type:varchar(255);collate:utf8mb4_unicode_ci" json:"issuer_url_service" form:"issuer_url_service"`
	IssuerHost       string    `gorm:"column:issuer_host;type:varchar(255);collate:utf8mb4_unicode_ci" json:"issuer_host" form:"issuer_host"`
	Status           int       `gorm:"column:status;default:0" json:"status" form:"status"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy        int       `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy        int       `gorm:"column:updated_by" json:"updated_by"`
}

func (Issuer) TableName() string {
	return "issuer"
}
