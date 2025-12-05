package modelOdoo

import (
	"time"
)

type IIDTransactionJournal struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Code        string    `gorm:"column:code" json:"code"`
	Description string    `gorm:"column:description" json:"description"`
	SecretKey   string    `gorm:"column:secret_key" json:"secret_key"`
	CreateDate  time.Time `gorm:"column:create_date;autoCreateTime" json:"create_date"`
	WriteDate   time.Time `gorm:"column:write_date;autoUpdateTime" json:"write_date"`

	// JournalID    int64           `gorm:"column:journal_id" json:"journal_id"`
	// ProductID    sql.NullInt64   `gorm:"column:product_id" json:"product_id"`
	// CreateUID    sql.NullInt64   `gorm:"column:create_uid" json:"create_uid"`
	// WriteUID     sql.NullInt64   `gorm:"column:write_uid" json:"write_uid"`
	// AccountID    int64           `gorm:"column:account_id" json:"account_id"`
	// MdrDefaultID sql.NullInt64   `gorm:"column:mdr_default_id" json:"mdr_default_id"`
	// CurrencyID   sql.NullInt64   `gorm:"column:currency_id" json:"currency_id"`
	// MaxTrxYear   sql.NullFloat64 `gorm:"column:max_trx_year" json:"max_trx_year"`
	// MaxTrx       sql.NullFloat64 `gorm:"column:max_trx" json:"max_trx"`
	// PartnerID    sql.NullInt64   `gorm:"column:partner_id" json:"partner_id"`
}

// TableName overrides the default table name
func (IIDTransactionJournal) TableName() string {
	return "iid_transaction_journal"
}
