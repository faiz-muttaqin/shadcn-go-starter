package modelOdoo

import (
	"database/sql"
)

type ResPartnerBank struct {
	ID         int64         `gorm:"column:id;primaryKey;autoIncrement:false"`
	PartnerID  int64         `gorm:"column:partner_id;not null;index:res_partner_bank__partner_id_index"`
	BankID     sql.NullInt64 `gorm:"column:bank_id;index;default:null"`
	Sequence   sql.NullInt32 `gorm:"column:sequence;default:null"`
	CurrencyID sql.NullInt64 `gorm:"column:currency_id;index;default:null"`
	CompanyID  sql.NullInt64 `gorm:"column:company_id;index;default:null"`
	CreateUID  sql.NullInt64 `gorm:"column:create_uid;index;default:null"`
	WriteUID   sql.NullInt64 `gorm:"column:write_uid;index;default:null"`

	AccNumber          string `gorm:"column:acc_number;type:varchar;not null"`
	SanitizedAccNumber string `gorm:"column:sanitized_acc_number;type:varchar;uniqueIndex:res_partner_bank_unique_number;default:null"`
	AccHolderName      string `gorm:"column:acc_holder_name;type:varchar;default:null"`

	Active                  sql.NullBool `gorm:"column:active;default:null"`
	AllowOutPayment         sql.NullBool `gorm:"column:allow_out_payment;default:null"`
	HasIbanWarning          sql.NullBool `gorm:"column:has_iban_warning;default:null"`
	HasMoneyTransferWarning sql.NullBool `gorm:"column:has_money_transfer_warning;default:null"`

	CreateDate sql.NullTime `gorm:"column:create_date;default:null"`
	WriteDate  sql.NullTime `gorm:"column:write_date;default:null"`

	// Relations (optional – bisa diaktifkan kalau model terkait sudah ada)
	// Partner    *ResPartner  `gorm:"foreignKey:PartnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Bank       *ResBank     `gorm:"foreignKey:BankID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:SET NULL"`
	// Currency   *ResCurrency `gorm:"foreignKey:CurrencyID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:SET NULL"`
	// Company    *ResCompany  `gorm:"foreignKey:CompanyID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:SET NULL"`
	// CreateUser *ResUser     `gorm:"foreignKey:CreateUID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:SET NULL"`
	// WriteUser  *ResUser     `gorm:"foreignKey:WriteUID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:SET NULL"`
}

func (ResPartnerBank) TableName() string {
	return "res_partner_bank"
}
