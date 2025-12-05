package modelParam

type ResponseCodeTrx struct {
	ID          int    `gorm:"primaryKey;autoIncrement;column:id" json:"id" form:"id"`
	Code        string `gorm:"unique;column:code" json:"code" form:"code"`
	Description string `gorm:"column:description" json:"description" form:"description"`
}

func (ResponseCodeTrx) TableName() string {
	return "response_code_trx"
}
