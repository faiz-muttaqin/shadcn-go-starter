package modelTrx

import (
	"time"
)

type SettlementDetail struct {
	ID                int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SettlementID      string    `gorm:"column:settlement_id" json:"settlement_id"` //edit by aziz
	TransactionID     string    `gorm:"column:transaction_id;size:255" json:"transaction_id"`
	TransactionType   string    `gorm:"column:transaction_type;size:10" json:"transaction_type"`
	MID               string    `gorm:"column:mid;size:255" json:"mid"`
	TID               string    `gorm:"column:tid;size:255" json:"tid"`
	CardType          string    `gorm:"column:card_type;size:255" json:"card_type"`
	Amount            int       `gorm:"column:amount" json:"amount"`
	TransactionDate   time.Time `gorm:"column:transaction_date" json:"transaction_date"`
	Trace             string    `gorm:"column:trace;size:255" json:"trace"`
	ResponseCode      string    `gorm:"column:response_code;size:255" json:"response_code"`
	ResponseAt        time.Time `gorm:"column:response_at" json:"response_at"`
	ApprovalCode      string    `gorm:"column:approval_code;size:255" json:"approval_code"`
	ReffID            string    `gorm:"column:reff_id;size:255" json:"reff_id"`
	IssuerID          int       `gorm:"column:issuer_id" json:"issuer_id"`
	VisaTID           string    `gorm:"column:visa_tid;size:255" json:"visa_tid"`
	VisaProductID     int       `gorm:"column:visa_product_id" json:"visa_product_id"`
	MasterNetworkData string    `gorm:"column:master_network_data;size:255" json:"master_network_data"`
	ContactlessFlag   int       `gorm:"column:contactless_flag" json:"contactless_flag"`
	PFID              int       `gorm:"column:pfid" json:"pfid"`
	VoidID            string    `gorm:"column:void_id;size:255" json:"void_id"`
	PAN               string    `form:"pan" json:"pan" gorm:"column:pan;size:255"`
	BankCode          string    `gorm:"size:255;column:bank_code" json:"bank_code"`
	Status            int       `gorm:"column:status" json:"status"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
}

func (SettlementDetail) TableName() string {
	return "settlement_details"
}
