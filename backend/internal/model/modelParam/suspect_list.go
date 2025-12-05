package modelParam

type SuspectList struct {
	ID     int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MID    string `gorm:"column:mid;type:mediumtext;not null" json:"mid" form:"mid"`
	TID    string `gorm:"column:tid;type:mediumtext;not null" json:"tid" form:"tid"`
	Trace  string `gorm:"column:trace;type:mediumtext;not null" json:"trace" form:"trace"`
	PAN    string `gorm:"column:pan;type:mediumtext;not null" json:"pan" form:"pan"`
	Date   string `gorm:"column:date;type:mediumtext;not null" json:"date" form:"date"`
	Status string `gorm:"column:status;type:mediumtext;not null" json:"status" form:"status"`
	Data   string `gorm:"column:data;type:mediumtext" json:"data" form:"data"`
}

// TableName sets the insert table name for this struct type
func (SuspectList) TableName() string {
	return "suspect_list"
}
