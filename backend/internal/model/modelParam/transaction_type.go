package modelParam

import "time"

type TransactionTypes struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Code      string    `gorm:"column:code;type:varchar(255);collate:utf8mb4_unicode_ci" json:"code" form:"code"`
	Name      string    `gorm:"column:name;type:varchar(255);collate:utf8mb4_unicode_ci" json:"name" form:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by"`
}

func (TransactionTypes) TableName() string {
	return "transaction_types"
}
