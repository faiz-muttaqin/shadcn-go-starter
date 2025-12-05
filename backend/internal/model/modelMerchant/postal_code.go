package modelMerchant

type PostalCodeID struct {
	ID          uint   `gorm:"primarykey"`
	Province    string `json:"province" gorm:"column:province;size:100"`
	City        string `json:"city" gorm:"column:city;size:100"`
	District    string `json:"district" gorm:"column:district;size:100"`
	Subdistrict string `json:"subdistrict" gorm:"column:subdistrict;size:100"`
	PostalCode  string `json:"postal_code" gorm:"column:postal_code;size:10"`
}

func (PostalCodeID) TableName() string {
	return "postal_code_id"
}
