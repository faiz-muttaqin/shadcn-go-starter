package modelMerchant

import "time"

type MerchantCategoryCode struct {
	ID             uint      `gorm:"primaryKey;column:id" json:"id"`
	MCC            string    `gorm:"size:4;unique;not null;column:mcc" json:"mcc"`                    // MCC code, e.g., "5812"
	IndonesianName string    `gorm:"type:varchar(255);column:name_indonesian" json:"name_indonesian"` // Category name in Indonesian
	EnglishName    string    `gorm:"type:varchar(255);column:name_english" json:"name_english"`       // Category name in English
	Description    string    `gorm:"type:text;column:description" json:"description,omitempty"`       // Optional: MCC description
	CreatedAt      time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
}

// TableName overrides the default table name used by GORM
func (MerchantCategoryCode) TableName() string {
	return "merchant_category_codes"
}
