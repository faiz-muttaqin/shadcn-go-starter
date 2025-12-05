package modelMerchant

import (
	"time"
)

type TerminatedMerchant struct {
	ID                   int64     `json:"id" gorm:"column:id;primarykey" form:"id"`
	NOMOR_PEMUTUSAN      string    `json:"NOMOR_PEMUTUSAN" gorm:"column:NOMOR_PEMUTUSAN;default:0" form:"NOMOR_PEMUTUSAN"`
	NAMA_MERCHANT        string    `json:"NAMA_MERCHANT" gorm:"column:NAMA_MERCHANT;size:200" form:"NAMA_MERCHANT"`
	MID                  string    `json:"MID" gorm:"column:MID;size:200" form:"MID"`
	EMAIL                string    `json:"EMAIL" gorm:"column:EMAIL;size:500" label:"EMAIL" form:"EMAIL"`
	KETERANGAN_PEMUTUSAN string    `json:"KETERANGAN_PEMUTUSAN" gorm:"column:KETERANGAN_PEMUTUSAN;size:100" label:"KETERANGAN_PEMUTUSAN" form:"KETERANGAN_PEMUTUSAN"`
	DIKETAHUI_OLEH       string    `json:"DIKETAHUI_OLEH" gorm:"column:DIKETAHUI_OLEH;size:100" label:"DIKETAHUI_OLEH" form:"DIKETAHUI_OLEH"`
	DISETUJUI_OLEH       string    `json:"DISETUJUI_OLEH" gorm:"column:DISETUJUI_OLEH;size:100" label:"DISETUJUI_OLEH" form:"DISETUJUI_OLEH"`
	DITANGANI_OLEH       string    `json:"DITANGANI_OLEH" gorm:"column:DITANGANI_OLEH;size:100" label:"DITANGANI_OLEH" form:"DITANGANI_OLEH"`
	DIBUAT_PADA          time.Time `json:"DIBUAT_PADA" gorm:"column:DIBUAT_PADA;autoCreateTime" form:"DIBUAT_PADA"`
	DIUBAH_PADA          time.Time `json:"DIUBAH_PADA" gorm:"column:DIUBAH_PADA;autoUpdateTime" form:"DIUBAH_PADA"`
}

func (TerminatedMerchant) TableName() string {
	return "terminated_merchant"
}
