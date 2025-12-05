package modelOdoo

import (
	"database/sql"
)

// ResBank merepresentasikan tabel res_bank di database Odoo.
type ResBank struct {
	ID           int64           `gorm:"primaryKey;autoIncrement:false;column:id"` // pakai sequence di DB
	State        sql.NullInt64   `gorm:"column:state;index"`                       // FK ke res_country_state.id
	Country      sql.NullInt64   `gorm:"column:country;index"`                     // FK ke res_country.id
	CreateUID    sql.NullInt64   `gorm:"column:create_uid;index"`                  // FK ke res_users.id
	WriteUID     sql.NullInt64   `gorm:"column:write_uid;index"`                   // FK ke res_users.id
	Name         string          `gorm:"column:name;not null"`                     // Nama bank
	Street       sql.NullString  `gorm:"column:street"`                            // Alamat baris 1
	Street2      sql.NullString  `gorm:"column:street2"`                           // Alamat baris 2
	Zip          sql.NullString  `gorm:"column:zip"`                               // Kode pos
	City         sql.NullString  `gorm:"column:city"`                              // Kota
	Email        sql.NullString  `gorm:"column:email"`                             // Email
	Phone        sql.NullString  `gorm:"column:phone"`                             // Telepon
	BIC          sql.NullString  `gorm:"column:bic;index:res_bank__bic_index"`     // Kode BIC/SWIFT
	Active       sql.NullBool    `gorm:"column:active"`                            // Status aktif
	CreateDate   sql.NullTime    `gorm:"column:create_date;autoCreateTime"`        // Tanggal dibuat
	WriteDate    sql.NullTime    `gorm:"column:write_date;autoUpdateTime"`         // Tanggal diubah
	JournalID    sql.NullInt64   `gorm:"column:journal_id;index"`                  // FK ke account_journal.id
	CurrencyID   sql.NullInt64   `gorm:"column:currency_id;index"`                 // FK ke res_currency.id
	TransferCost sql.NullFloat64 `gorm:"column:transfer_cost"`                     // Biaya transfer
}

// TableName menentukan nama tabel di database.
func (ResBank) TableName() string {
	return "res_bank"
}
