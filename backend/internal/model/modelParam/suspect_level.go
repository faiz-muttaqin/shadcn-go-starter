package modelParam

import (
	"time"
)

type SuspectLevel struct {
	ID           int       `gorm:"column:id;primaryKey" json:"id"`
	Name         string    `gorm:"column:name;size:255" json:"name"`
	CountToBlock int       `gorm:"column:count_to_block" json:"count_to_block"`
	BlockTime    int       `gorm:"column:block_time;comment:per second" json:"block_time"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UpdatedBy    string    `gorm:"column:updated_by;size:255" json:"updated_by"`
}

// TableName override
func (SuspectLevel) TableName() string {
	return "suspect_level"
}
