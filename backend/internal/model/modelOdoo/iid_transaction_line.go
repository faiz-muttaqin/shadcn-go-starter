package modelOdoo

import (
	"database/sql"
	"time"
)

type IidTransactionLine struct {
	ID             int             `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Time           string          `gorm:"column:time" json:"time"`
	Host           string          `gorm:"column:host" json:"host"`
	Note           string          `gorm:"column:note" json:"note"`
	Date           sql.NullTime    `gorm:"column:date" json:"date"`
	Amount         sql.NullFloat64 `gorm:"column:amount" json:"amount"`
	Datetime       sql.NullTime    `gorm:"column:datetime" json:"datetime"`
	CreateDate     time.Time       `gorm:"column:create_date;autoCreateTime" json:"create_date"`
	WriteDate      time.Time       `gorm:"column:write_date;autoUpdateTime" json:"write_date"`
	NoReference    string          `gorm:"column:no_reference" json:"no_reference"`
	NoRetrievalRef string          `gorm:"column:no_retrieval_ref" json:"no_retrieval_ref"`
	BatchRecord    int             `gorm:"column:batch_record" json:"batch_record"`
	IsBatchClear   bool            `gorm:"column:is_batch_clear" json:"is_batch_clear"`
	BatchRef       string          `gorm:"column:batch_ref" json:"batch_ref"`
	// PartnerID         int             `gorm:"column:partner_id" json:"partner_id"` // NMID
	// CurrencyID        int             `gorm:"column:currency_id" json:"currency_id"`
	// TransactionTypeID int             `gorm:"column:transaction_type_id" json:"transaction_type_id"`
	// OrderID           int             `gorm:"column:order_id" json:"order_id"`
	// OrderLineID       int             `gorm:"column:order_line_id" json:"order_line_id"`
	// ProductID         int             `gorm:"column:product_id" json:"product_id"`
	// CreateUID         int             `gorm:"column:create_uid" json:"create_uid"`
	// WriteUID          int             `gorm:"column:write_uid" json:"write_uid"`
	// Service           string          `gorm:"column:service" json:"service"`
}

// TableName override
func (IidTransactionLine) TableName() string {
	return "iid_transaction_line"
}
