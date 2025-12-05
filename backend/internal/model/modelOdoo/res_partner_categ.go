package modelOdoo

import (
	"database/sql"
	"time"
)

type ResPartnerCateg struct {
	ID         int64         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreateUID  sql.NullInt64 `gorm:"column:create_uid" json:"create_uid"`
	WriteUID   sql.NullInt64 `gorm:"column:write_uid" json:"write_uid"`
	Name       string        `gorm:"column:name;size:255" json:"name"`
	CreateDate time.Time     `gorm:"column:create_date;autoCreateTime" json:"create_date"`
	WriteDate  time.Time     `gorm:"column:write_date;autoUpdateTime" json:"write_date"`
}

// TableName override
func (ResPartnerCateg) TableName() string {
	return "res_partner_categ"
}
