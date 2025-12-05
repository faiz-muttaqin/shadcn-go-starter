package modelParam

type ResponseCodeReversal struct {
	ID          int    `gorm:"primaryKey;autoIncrement;column:id" json:"id" form:"id"`
	Code        string `gorm:"unique;column:code" json:"code" form:"code"`
	Description string `gorm:"column:description" json:"description" form:"description"`
}

func (ResponseCodeReversal) TableName() string {
	return "response_code_reversal"
}
