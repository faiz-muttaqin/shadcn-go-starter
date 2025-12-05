package modelMerchant

import "time"

type Terminal struct {
	TerminalId        int64     `json:"id" form:"id" gorm:"column:id;primarykey"`
	MerchantId        int64     `json:"merchant_id" form:"merchant_id" gorm:"column:merchant_id"`
	Status            string    `json:"status" form:"status" gorm:"column:status;type:varchar(255)"`
	LimitAmt          string    `json:"limit_amt" form:"limit_amt" gorm:"column:limit_amt;size:30"`
	SerialNumber      string    `json:"serial_number" form:"serial_number" gorm:"column:serial_number;type:varchar(255)"`
	IMEI              string    `json:"imei" form:"imei" gorm:"column:imei;type:varchar(255)"`
	CreatedAt         time.Time `json:"created_at" form:"created_at" gorm:"column:created_at"`
	UpdatedAt         time.Time `json:"updated_at" form:"updated_at" gorm:"column:updated_at"`
	Email             string    `json:"email" form:"email" gorm:"column:email;type:varchar(255)"`
	Secret            string    `json:"-" form:"-" gorm:"column:secret;type:varchar(255)"`
	Password          string    `json:"-" form:"-" gorm:"column:password;type:varchar(255)"`
	PasswordTemp      string    `json:"-" form:"-" gorm:"column:password_temp;type:varchar(255)"`
	IsEmailSended     bool      `json:"-" form:"-" gorm:"column:is_email_sended"`
	IsPasswordRenewed bool      `json:"-" form:"-" gorm:"column:is_password_renewed"`
	LastLogin         int64     `json:"-" form:"-" gorm:"column:last_login"`
	Session           string    `json:"-" form:"-" gorm:"column:session;type:varchar(255)"`
	SessionExpired    int64     `json:"-" form:"-" gorm:"column:session_expired"`
}

func (Terminal) TableName() string {
	return "terminals"
}
