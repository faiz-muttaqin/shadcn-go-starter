package modelParam

import "time"

type TransactionStatus struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Status    string    `gorm:"column:status;type:varchar(255);collate:utf8mb4_unicode_ci" json:"status" form:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by"`
}

func (TransactionStatus) TableName() string {
	return "transaction_status"
}
