package modelParam

type ParamFDS struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Api        string `gorm:"column:api;size:30" json:"api" form:"api"`
	Value      string `gorm:"column:value;size:20" json:"value" form:"value"`
	Keterangan string `gorm:"column:keterangan;size:50" json:"keterangan" form:"keterangan"`
}

func (ParamFDS) TableName() string {
	return "param_fds"
}
