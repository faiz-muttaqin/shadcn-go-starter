package modelOdoo

import "time"

type ListAgent struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreateUID   int       `gorm:"column:create_uid" json:"create_uid"`
	WriteUID    int       `gorm:"column:write_uid" json:"write_uid"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	CreateDate  time.Time `gorm:"column:create_date;autoCreateTime" json:"create_date"`
	WriteDate   time.Time `gorm:"column:write_date;autoUpdateTime" json:"write_date"`
	DisplayName string    `gorm:"column:display_name" json:"display_name"`
}

// TableName overrides the default table name.
func (ListAgent) TableName() string {
	return "list_agent"
}
