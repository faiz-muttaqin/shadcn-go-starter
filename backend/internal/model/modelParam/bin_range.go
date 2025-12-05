package modelParam

type BinRange struct {
	ID           int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	BankCode     string `gorm:"column:bank_code;type:varchar(255)" json:"bank_code" form:"bank_code"`
	CardType     string `gorm:"column:card_type;type:varchar(255)" json:"card_type" form:"card_type"`
	Name         string `gorm:"column:name;type:varchar(255);collate:utf8mb4_unicode_ci" json:"name" form:"name"`
	PanRangeLow  string `gorm:"column:pan_range_low;type:varchar(255);collate:utf8mb4_unicode_ci" json:"pan_range_low" form:"pan_range_low"`
	PanRangeHigh string `gorm:"column:pan_range_high;type:varchar(255);collate:utf8mb4_unicode_ci" json:"pan_range_high" form:"pan_range_high"`
	IssuerID     int    `gorm:"column:issuer_id" json:"issuer_id"`
}

func (BinRange) TableName() string {
	return "bin_range"
}
