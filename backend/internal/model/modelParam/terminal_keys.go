package modelParam

import "time"

type TerminalKeys struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TID       string    `gorm:"column:tid" json:"tid" form:"tid"`
	KeyType   string    `gorm:"column:key_type" json:"key_type" form:"key_type"`
	Value     string    `gorm:"column:value" json:"value" form:"value"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (TerminalKeys) TableName() string {
	return "terminal_keys"
}
