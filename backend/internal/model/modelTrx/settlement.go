package modelTrx

import (
	"time"
)

type Settlement struct {
	ID               int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SettlementID     string    `gorm:"column:settlement_id" json:"settlement_id"` //add by aziz
	MID              string    `gorm:"column:mid" json:"mid"`
	TID              string    `gorm:"column:tid" json:"tid"`
	Trace            string    `gorm:"column:trace" json:"trace"`
	Batch            string    `gorm:"column:batch" json:"batch"`
	CurrencyCode     string    `gorm:"column:currency_code;default:'360'" json:"currency_code"`
	BankID           int       `gorm:"column:bank_id" json:"bank_id"`
	FirstTrxTime     time.Time `gorm:"column:first_trx_time" json:"first_trx_time"`
	SubBatchNo       string    `gorm:"column:sub_batch_no" json:"sub_batch_no"`
	SettleDate       time.Time `gorm:"column:settle_date" json:"settle_date"`
	TotalTransaction int       `gorm:"column:total_transaction" json:"total_transaction"`
	TotalAmount      int       `gorm:"column:total_amount" json:"total_amount"`
	HostSaleCount    int       `gorm:"column:host_sale_count" json:"host_sale_count"`
	HostSaleAmount   int       `gorm:"column:host_sale_amount" json:"host_sale_amount"`
	HostRefundCount  int       `gorm:"column:host_refund_count" json:"host_refund_count"`
	HostRefundAmount int       `gorm:"column:host_refund_amount" json:"host_refund_amount"`
	PosSaleCount     int       `gorm:"column:pos_sale_count" json:"-"`
	PosSaleAmount    int       `gorm:"column:pos_sale_amount" json:"-"`
	PosRefundCount   int       `gorm:"column:pos_refund_count" json:"-"`
	PosRefundAmount  int       `gorm:"column:pos_refund_amount" json:"-"`
	ClearingName     string    `gorm:"column:clearing_name" json:"-"`
	ClearingStatus   string    `gorm:"column:clearing_status" json:"-"`
	ClearingFlag     int       `gorm:"column:clearing_flag" json:"-"`
	ClearingDate     time.Time `gorm:"column:clearing_date" json:"-"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"` // Not Voewed
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"` // Not Voewed
}

func (Settlement) TableName() string {
	return "settlement"
}
