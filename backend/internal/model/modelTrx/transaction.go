package modelTrx

import (
	"time"
)

//	type Transaction struct {
//		ID                int       `form:"id" json:"id" gorm:"column:id;primaryKey;autoIncrement"`
//		TransactionID     string    `form:"transaction_id" json:"transaction_id" gorm:"column:transaction_id;size:255;not null"`
//		TransactionType   string    `form:"transaction_type" json:"transaction_type" gorm:"column:transaction_type;size:10"`
//		MID               string    `form:"mid" json:"mid" gorm:"column:mid;size:255"`
//		TID               string    `form:"tid" json:"tid" gorm:"column:tid;size:255"`
//		Procode           string    `form:"procode" json:"procode" gorm:"column:procode;size:255"`
//		RRN               string    `form:"rrn" json:"rrn" gorm:"column:rrn;size:255"`
//		CardType          string    `form:"card_type" json:"card_type" gorm:"column:card_type;size:255"`
//		Batch             string    `form:"batch" json:"batch" gorm:"column:batch"`
//		TransMode         string    `form:"trans_mode" json:"trans_mode" gorm:"column:trans_mode"`
//		PAN               string    `form:"pan" json:"pan" gorm:"column:pan;size:255"`
//		PANEnc            string    `form:"-" json:"-" gorm:"column:pan_enc;size:255"`
//		TrackData         string    `form:"-" json:"-" gorm:"column:track_data;size:255"`
//		EMVTag            string    `form:"-" json:"-" gorm:"column:emv_tag;type:text"`
//		Amount            int       `form:"amount" json:"amount" gorm:"column:amount"`
//		TransactionDate   time.Time `json:"transaction_date" gorm:"column:transaction_date"`
//		ResponseAt        time.Time `json:"response_at" gorm:"column:response_at"`
//		CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
//		UpdatedAt         time.Time `json:"-" gorm:"column:updated_at"`
//		SettledAt         time.Time `json:"settled_at" gorm:"column:settled_at"`
//		Trace             string    `form:"trace" json:"trace" gorm:"column:trace;size:255"`
//		ISORequest        string    `form:"-" json:"-" gorm:"column:iso_request;type:text"`
//		VisaTID           string    `form:"-" json:"-" gorm:"column:visa_tid;size:255"`
//		VisaProductID     int       `form:"-" json:"-" gorm:"column:visa_product_id"`
//		MasterNetworkData string    `form:"-" json:"-" gorm:"column:master_network_data;size:255"`
//		ContactlessFlag   int       `form:"-" json:"-" gorm:"column:contactless_flag"`
//		PFID              int       `form:"-" json:"-" gorm:"column:pfid"`
//		ResponseCode      string    `form:"response_code" json:"response_code" gorm:"column:response_code;size:255"`
//		ISOResponse       string    `form:"-" json:"-" gorm:"column:iso_response;type:text"`
//		ApprovalCode      string    `form:"approval_code" json:"approval_code" gorm:"column:approval_code;size:255"`
//		ReffID            string    `form:"-" json:"-" gorm:"column:reff_id;size:255"`
//		IssuerID          int       `form:"issuer_id" json:"issuer_id" gorm:"column:issuer_id"`
//		Status            int       `form:"status" json:"status" gorm:"column:status"`
//		Longitude         string    `form:"longitude" json:"longitude" gorm:"column:longitude;size:255"`
//		Latitude          string    `form:"latitude" json:"latitude" gorm:"column:latitude;size:255"`
//		VoidID            string    `form:"-" json:"-" gorm:"column:void_id;size:255"`
//		SettleFlag        int       `form:"-" json:"-" gorm:"column:settle_flag;default:0"`
//		ReversalFlag      int       `form:"-" json:"-" gorm:"column:reversal_flag;default:0"`
//		BatchUFlag        int       `form:"-" json:"-" gorm:"column:batch_u_flag;default:1"`
//	}
type Transaction struct {
	ID              int       `form:"id" json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	TransactionID   string    `form:"transaction_id" json:"transaction_id" gorm:"column:transaction_id;size:255;not null"`
	TransactionType string    `form:"transaction_type" json:"transaction_type" gorm:"column:transaction_type;size:10"`
	TransMode       string    `form:"trans_mode" json:"trans_mode" gorm:"column:trans_mode"`
	Procode         string    `form:"procode" json:"procode" gorm:"column:procode;size:255"`
	MID             string    `form:"mid" json:"mid" gorm:"column:mid;size:255"`
	TID             string    `form:"tid" json:"tid" gorm:"column:tid;size:255"`
	RRN             string    `form:"rrn" json:"rrn" gorm:"column:rrn;size:255"`
	Batch           string    `form:"batch" json:"batch" gorm:"column:batch"`
	Trace           string    `form:"trace" json:"trace" gorm:"column:trace;size:255"`
	PAN             string    `form:"pan" json:"pan" gorm:"column:pan;size:255"`
	Amount          int       `form:"amount" json:"amount" gorm:"column:amount"`
	TransactionDate time.Time `json:"transaction_date" gorm:"column:transaction_date"`
	ResponseCode    string    `form:"response_code" json:"response_code" gorm:"column:response_code;size:255"`
	ApprovalCode    string    `form:"approval_code" json:"approval_code" gorm:"column:approval_code;size:255"`
	IssuerID        int       `form:"issuer_id" json:"issuer_id" gorm:"column:issuer_id"`
	Signature       string    `form:"signature" json:"signature" gorm:"column:signature"`
	Longitude       string    `form:"longitude" json:"longitude" gorm:"column:longitude;size:255"`
	Latitude        string    `form:"latitude" json:"latitude" gorm:"column:latitude;size:255"`
	CardType        string    `form:"card_type" json:"card_type" gorm:"column:card_type;size:255"`
	ResponseAt      time.Time `json:"response_at" gorm:"column:response_at"`
	SettledAt       time.Time `json:"-" gorm:"column:settled_at"`
	CreatedAt       time.Time `json:"-" gorm:"column:created_at"`
	Status          int       `form:"-" json:"-" gorm:"column:status"`
	UpdatedAt       time.Time `json:"-" gorm:"column:updated_at"`
	PANEnc          string    `form:"-" json:"-" gorm:"column:pan_enc;size:255"`
	EMVTag          string    `form:"-" json:"-" gorm:"column:emv_tag;type:text"`
	ISORequest      string    `form:"-" json:"-" gorm:"column:iso_request;type:text"`
	// VisaTID           string    `form:"-" json:"-" gorm:"column:visa_tid;size:255"`
	// VisaProductID     int       `form:"-" json:"-" gorm:"column:visa_product_id"`
	MasterNetworkData string `form:"-" json:"-" gorm:"column:master_network_data;size:255"`
	ContactlessFlag   int    `form:"-" json:"-" gorm:"column:contactless_flag"`
	PFID              int    `form:"-" json:"-" gorm:"column:pfid"`
	ReffID            string `form:"-" json:"-" gorm:"column:reff_id;size:255"`
	VoidID            string `form:"-" json:"-" gorm:"column:void_id;size:255"`
	SettleFlag        int    `form:"-" json:"-" gorm:"column:settle_flag;default:0"`
	ReversalFlag      int    `form:"-" json:"-" gorm:"column:reversal_flag;default:0"`
	BatchUFlag        int    `form:"-" json:"-" gorm:"column:batch_u_flag;default:1"`
	// TrackData       string    `form:"-" json:"-" gorm:"column:track_data;size:255"`
	// ISOResponse       string `form:"-" json:"-" gorm:"column:iso_response;type:text"`
}

func (Transaction) TableName() string {
	return "transactions"
}
