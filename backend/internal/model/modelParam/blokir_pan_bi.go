package modelParam

// BlokirPanBI mewakili struktur tabel `blokir_pan_bi` dalam database.
type BlokirPanBI struct {
	ID    uint   `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	NoPan string `gorm:"type:varchar(25);not null;column:no_pan" json:"no_pan"`
}

// TableName mendefinisikan nama tabel yang akan digunakan GORM.
func (BlokirPanBI) TableName() string {
	return "blokir_pan_bi"
}
