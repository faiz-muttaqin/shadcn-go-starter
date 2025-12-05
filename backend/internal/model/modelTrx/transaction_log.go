package modelTrx

import (
	"time"
)

//	type TransactionsLog struct {
//		ID                 int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
//		TransactionID      string    `gorm:"column:transaction_id;size:255;not null" json:"transaction_id"`
//		TransactionType    string    `gorm:"column:transaction_type;size:10" json:"transaction_type"`
//		MID                string    `gorm:"column:mid;size:255" json:"mid"`
//		TID                string    `gorm:"column:tid;size:255" json:"tid"`
//		CardType           string    `gorm:"column:card_type;size:255" json:"card_type"`
//		PAN                string    `gorm:"column:pan;size:255" json:"-"`
//		PANEnc             string    `gorm:"column:pan_enc;size:255" json:"-"`
//		TrackData          string    `gorm:"column:track_data;size:255" json:"track_data"`
//		EMVTag             string    `gorm:"column:emv_tag;type:text" json:"-"`
//		Amount             int       `gorm:"column:amount" json:"amount"`
//		TransactionDate    time.Time `gorm:"column:transaction_date" json:"-"`
//		TransactionDateStr string    `gorm:"-" json:"transaction_date_str"`
//		Trace              string    `gorm:"column:trace;size:255" json:"trace"`
//		ISORequest         string    `gorm:"column:iso_request;type:text" json:"-"`
//		VisaTID            string    `gorm:"column:visa_tid;size:255" json:"-"`
//		VisaProductID      int       `gorm:"column:visa_product_id" json:"-"`
//		MasterNetworkData  string    `gorm:"column:master_network_data;size:255" json:"-"`
//		ContactlessFlag    int       `gorm:"column:contactless_flag" json:"contactless_flag"`
//		PFID               int       `gorm:"column:pfid" json:"pfid"`
//		ResponseCode       string    `gorm:"column:response_code;size:255" json:"response_code"`
//		ResponseAt         time.Time `gorm:"column:response_at" json:"-"`
//		ResponseAtStr      string    `gorm:"-" json:"response_at_str"`
//		ISOResponse        string    `gorm:"column:iso_response;type:text" json:"-"`
//		ApprovalCode       string    `gorm:"column:approval_code;size:255" json:"approval_code"`
//		ReffID             string    `gorm:"column:reff_id;size:255" json:"reff_id"`
//		IssuerID           int       `gorm:"column:issuer_id" json:"issuer_id"`
//		IssuerName         string    `gorm:"-" json:"issuer_name"`
//		Status             int       `gorm:"column:status" json:"status"`
//		StatusStr          string    `gorm:"-" json:"status_str"`
//		Longitude          string    `gorm:"column:longitude;size:255" json:"longitude"`
//		Latitude           string    `gorm:"column:latitude;size:255" json:"latitude"`
//		VoidID             string    `gorm:"column:void_id;size:255" json:"void_id"`
//		SettleFlag         int       `gorm:"column:settle_flag;default:0" json:"settle_flag"`
//		ReversalFlag       int       `gorm:"column:reversal_flag;default:0" json:"reversal_flag"`
//		CreatedAt          time.Time `gorm:"column:created_at" json:"-"`
//		CreatedAtStr       string    `gorm:"-" json:"created_at_str"`
//		UpdatedAt          time.Time `gorm:"column:updated_at" json:"-"`
//		UpdatedAtStr       string    `gorm:"-" json:"updated_at_str"`
//		SettledAt          time.Time `gorm:"column:settled_at" json:"-"`
//		SettledAtStr       string    `gorm:"-" json:"settled_at_str"`
//		BatchUFlag         int       `gorm:"column:batch_u_flag;default:1" json:"batch_u_flag"`
//		LoggedAt           time.Time `gorm:"column:logged_at" json:"-"`
//		LoggedAtStr        string    `gorm:"-" json:"logged_at_str"`
//	}

type TransactionsLog struct {
	ID                int       `form:"id" json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	TransactionID     string    `form:"transaction_id" json:"transaction_id" gorm:"column:transaction_id;size:255;not null"`
	TransactionType   string    `form:"transaction_type" json:"transaction_type" gorm:"column:transaction_type;size:10"`
	TransMode         string    `form:"trans_mode" json:"trans_mode" gorm:"column:trans_mode"`
	Procode           string    `form:"procode" json:"procode" gorm:"column:procode;size:255"`
	MID               string    `form:"mid" json:"mid" gorm:"column:mid;size:255"`
	TID               string    `form:"tid" json:"tid" gorm:"column:tid;size:255"`
	RRN               string    `form:"rrn" json:"rrn" gorm:"column:rrn;size:255"`
	Batch             string    `form:"batch" json:"batch" gorm:"column:batch"`
	Trace             string    `form:"trace" json:"trace" gorm:"column:trace;size:255"`
	PAN               string    `form:"pan" json:"pan" gorm:"column:pan;size:255"`
	Amount            int       `form:"amount" json:"amount" gorm:"column:amount"`
	TransactionDate   time.Time `json:"transaction_date" gorm:"column:transaction_date"`
	ResponseCode      string    `form:"response_code" json:"response_code" gorm:"column:response_code;size:255"`
	ApprovalCode      string    `form:"approval_code" json:"approval_code" gorm:"column:approval_code;size:255"`
	IssuerID          int       `form:"issuer_id" json:"issuer_id" gorm:"column:issuer_id"`
	Signature         string    `form:"signature" json:"signature" gorm:"column:signature"`
	SettledAt         time.Time `json:"settled_at" gorm:"column:settled_at"`
	Longitude         string    `form:"longitude" json:"longitude" gorm:"column:longitude;size:255"`
	Latitude          string    `form:"latitude" json:"latitude" gorm:"column:latitude;size:255"`
	CardType          string    `form:"card_type" json:"card_type" gorm:"column:card_type;size:255"`
	ResponseAt        time.Time `json:"response_at" gorm:"column:response_at"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         time.Time `json:"-" gorm:"column:updated_at"`
	Status            int       `form:"-" json:"-" gorm:"column:status"`
	PANEnc            string    `form:"-" json:"-" gorm:"column:pan_enc;size:255"`
	EMVTag            string    `form:"-" json:"-" gorm:"column:emv_tag;type:text"`
	ISORequest        string    `form:"-" json:"-" gorm:"column:iso_request;type:text"`
	MasterNetworkData string    `form:"-" json:"-" gorm:"column:master_network_data;size:255"`
	ContactlessFlag   int       `form:"-" json:"-" gorm:"column:contactless_flag"`
	PFID              int       `form:"-" json:"-" gorm:"column:pfid"`
	VoidID            string    `form:"-" json:"-" gorm:"column:void_id;size:255"`
	SettleFlag        int       `form:"-" json:"-" gorm:"column:settle_flag;default:0"`
	ReversalFlag      int       `form:"-" json:"-" gorm:"column:reversal_flag;default:0"`
	BatchUFlag        int       `form:"-" json:"-" gorm:"column:batch_u_flag;default:1"`
	// ISOResponse       string    `form:"-" json:"-" gorm:"column:iso_response;type:text"`
	// ReffID            string    `form:"-" json:"-" gorm:"column:reff_id;size:255"`
	// TrackData         string    `form:"-" json:"-" gorm:"column:track_data;size:255"`
}

func (TransactionsLog) TableName() string {
	return "transaction_logs"
}
