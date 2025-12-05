package modelMerchant

import "time"

type LogActivity struct {
	ID        uint      `json:"id" gorm:"column:id;primarykey"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	MID       uint      `json:"mid" gorm:"column:mid"`
	FullName  string    `json:"full_name" gorm:"column:full_name"`
	Action    string    `json:"action" gorm:"column:action"`
	Status    string    `json:"status" gorm:"column:status"`
	Log       string    `json:"log" gorm:"column:log"`
	UserAgent string    `json:"user_agent" gorm:"column:user_agent"`
	ReqMethod string    `json:"req_method" gorm:"column:req_method"`
	IP        string    `json:"ip" gorm:"column:ip"`
	ReqUri    string    `json:"req_uri" gorm:"column:req_uri"`
}

func (LogActivity) TableName() string {
	return "log_activities"
}
