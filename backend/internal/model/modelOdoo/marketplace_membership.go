package modelOdoo

import (
	"database/sql"
	"time"
)

type MarketplaceMembership struct {
	ID           int64           `gorm:"column:id;primaryKey" json:"id"`
	MdrID        sql.NullInt64   `gorm:"column:mdr_id" json:"mdr_id"`
	FeeID        sql.NullInt64   `gorm:"column:fee_id" json:"fee_id"`
	CurrencyID   sql.NullInt64   `gorm:"column:currency_id" json:"currency_id"`
	CreateUID    sql.NullInt64   `gorm:"column:create_uid" json:"create_uid"`
	WriteUID     sql.NullInt64   `gorm:"column:write_uid" json:"write_uid"`
	Name         string          `gorm:"column:name" json:"name"`
	Code         string          `gorm:"column:code" json:"code"`
	FeeCost      sql.NullFloat64 `gorm:"column:fee_cost" json:"fee_cost"`
	MinTrx       sql.NullFloat64 `gorm:"column:min_trx" json:"min_trx"`
	MaxTrx       sql.NullFloat64 `gorm:"column:max_trx" json:"max_trx"`
	CreateDate   time.Time       `gorm:"column:create_date;autoCreateTime" json:"create_date"`
	WriteDate    time.Time       `gorm:"column:write_date;autoUpdateTime" json:"write_date"`
	JournalMdrID sql.NullInt64   `gorm:"column:journal_mdr_id" json:"journal_mdr_id"`
	JournalFeeID sql.NullInt64   `gorm:"column:journal_fee_id" json:"journal_fee_id"`
	MinBalance   sql.NullFloat64 `gorm:"column:min_balance" json:"min_balance"`
	MinWd        sql.NullFloat64 `gorm:"column:min_wd" json:"min_wd"`
	MntFee       sql.NullFloat64 `gorm:"column:mnt_fee" json:"mnt_fee"`
	PaymentDate  sql.NullTime    `gorm:"column:payment_date" json:"payment_date"`
}

// TableName overrides the default table name
func (MarketplaceMembership) TableName() string {
	return "marketplace_membership"
}
