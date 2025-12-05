package modelParam

type RouteFDS struct {
	ID         int    `json:"id" gorm:"column:id;primaryKey"`
	URL        string `json:"url" gorm:"column:url"`
	Keterangan string `json:"keterangan" gorm:"column:keterangan"`
	Status     int    `json:"status" gorm:"column:status;default:1"`
	Result     int    `json:"result" gorm:"column:result;default:2"`
	Data       string `json:"data" gorm:"column:data;default:'Suspect'"`
}

func (RouteFDS) TableName() string {
	return "route_fds"
}
