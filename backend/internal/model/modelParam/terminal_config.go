package modelParam

type TerminalConfig struct {
	ID    int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TID   string `gorm:"column:tid" json:"tid" form:"tid"`
	Batch string `gorm:"column:batch" json:"batch" form:"batch"`
	Trace string `gorm:"column:trace" json:"trace" form:"trace"`
}

func (TerminalConfig) TableName() string {
	return "terminal_config"
}
