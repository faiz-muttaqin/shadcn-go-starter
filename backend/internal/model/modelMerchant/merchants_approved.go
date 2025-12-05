package modelMerchant

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

var IGNORED_MERCHANT_APPROVED_KEY = []string{
	"PASSWORD_RAW",
	"PASSWORD",
	"DIUBAH_OLEH",
	"DIHAPUS_PADA",
	"KETERANGAN_PENDAFTARAN",
	"BENTUK_BADAN_HUKUM",
	"NIB",
	"DISTRICT",
	"SUBDISTRICT",
	// "BANK_PENERBIT_REKENING",
	// "NAMA_PEMILIK_REKENING",
	// "NOMOR_REKENING",
	"BANK_CODE",
	"NIK",
	"TEMPAT_LAHIR_PEMILIK_USAHA",
	"JENIS_KELAMIN_PEMILIK_USAHA",
	"PEKERJAAN_PEMILIK_USAHA",
	"FILE_NIB",
	"FILE_NPWP",
	"FILE_KTP",
	"FILE_SELFIE_KTP",
	"FILE_FOTO_DEPAN_USAHA",
	"FILE_SELFIE_LOKASI_USAHA",
	"DISETUJUI_OLEH",
	"DISETUJUI_PADA",
	"PENETAPAN_DEVICE_OLEH",
	"PENETAPAN_DEVICE_PADA",
	"DIUBAH_PADA",
}
var EDITABLE_MERCHANT_APPROVED_KEY = []string{
	"LOGIN_BLOCKED", "BANK_PENERBIT_REKENING", "NAMA_USAHA", "NAMA_PEMILIK_REKENING", "NOMOR_REKENING", "EMAIL", "ALAMAT_USAHA", "NOMOR_TELEPON_PEMILIK_USAHA",
}

type MerchantApproved struct {
	ID                             int64          `json:"id" gorm:"column:id;primarykey" form:"id"`
	STATUS_PENDAFTARAN             int            `json:"STATUS_PENDAFTARAN" gorm:"column:STATUS_PENDAFTARAN;default:0" form:"STATUS_PENDAFTARAN"`
	KETERANGAN_PENDAFTARAN         string         `json:"KETERANGAN_PENDAFTARAN" gorm:"column:KETERANGAN_PENDAFTARAN;size:1000" form:"KETERANGAN_PENDAFTARAN"`
	LOGIN_BLOCKED                  bool           `json:"LOGIN_BLOCKED" gorm:"column:LOGIN_BLOCKED" form:"LOGIN_BLOCKED"`
	NAMA_AGREGATOR                 string         `json:"NAMA_AGREGATOR" gorm:"column:NAMA_AGREGATOR;size:200" form:"NAMA_AGREGATOR"`
	NOMOR_PENDAFTARAN              string         `json:"NOMOR_PENDAFTARAN" gorm:"column:NOMOR_PENDAFTARAN;size:500" label:"NOMOR_PENDAFTARAN" form:"NOMOR_PENDAFTARAN"`
	EMAIL                          string         `json:"EMAIL" gorm:"column:EMAIL;size:500" label:"EMAIL" form:"EMAIL"`
	JENIS_PENDAFTARAN              string         `json:"JENIS_PENDAFTARAN" gorm:"column:JENIS_PENDAFTARAN;size:100" label:"JENIS_PENDAFTARAN" form:"JENIS_PENDAFTARAN"`
	BENTUK_BADAN_HUKUM             string         `json:"BENTUK_BADAN_HUKUM" gorm:"column:BENTUK_BADAN_HUKUM;size:100" label:"BENTUK_BADAN_HUKUM" form:"BENTUK_BADAN_HUKUM"`
	NIB                            string         `json:"NIB" gorm:"column:NIB;size:100" label:"NIB" form:"NIB"`
	NAMA_USAHA                     string         `json:"NAMA_USAHA" gorm:"column:NAMA_USAHA;size:500" label:"NAMA_USAHA" form:"NAMA_USAHA"`
	ALAMAT_USAHA                   string         `json:"ALAMAT_USAHA" gorm:"column:ALAMAT_USAHA;size:500" label:"ALAMAT_USAHA" form:"ALAMAT_USAHA"`
	KODE_POS_FULL                  string         `json:"KODE_POS_FULL" gorm:"column:KODE_POS_FULL;size:500" label:"KODE_POS_FULL" form:"KODE_POS_FULL"`
	KODE_POS                       string         `json:"KODE_POS" gorm:"column:KODE_POS;size:20" label:"KODE_POS" form:"KODE_POS"`
	PROVINCE                       string         `json:"PROVINCE" gorm:"column:PROVINCE;size:100" label:"PROVINSI" form:"PROVINCE"`
	CITY                           string         `json:"CITY" gorm:"column:CITY;size:100" label:"KOTA/KABUPATEN" form:"CITY"`
	DISTRICT                       string         `json:"DISTRICT" gorm:"column:DISTRICT;size:100" label:"KECAMATAN" form:"DISTRICT"`
	SUBDISTRICT                    string         `json:"SUBDISTRICT" gorm:"column:SUBDISTRICT;size:100" label:"DESA/KELURAHAN" form:"SUBDISTRICT"`
	NOMOR_TELEPON_USAHA            string         `json:"NOMOR_TELEPON_USAHA" gorm:"column:NOMOR_TELEPON_USAHA;size:20" label:"NOMOR_TELEPON_USAHA" form:"NOMOR_TELEPON_USAHA"`
	NAMA_PEMILIK_REKENING          string         `json:"NAMA_PEMILIK_REKENING" gorm:"column:NAMA_PEMILIK_REKENING;size:50" label:"NAMA_PEMILIK_REKENING" form:"NAMA_PEMILIK_REKENING"`
	NOMOR_REKENING                 string         `json:"NOMOR_REKENING" gorm:"column:NOMOR_REKENING;size:50" label:"NOMOR_REKENING" form:"NOMOR_REKENING"`
	BANK_CODE                      string         `json:"BANK_CODE" gorm:"column:BANK_CODE;size:50" label:"BANK_CODE" form:"BANK_CODE"`
	BANK_PENERBIT_REKENING         string         `json:"BANK_PENERBIT_REKENING" gorm:"column:BANK_PENERBIT_REKENING;size:50" label:"BANK_PENERBIT_REKENING" form:"BANK_PENERBIT_REKENING"`
	NIK                            string         `json:"NIK" gorm:"column:NIK;size:50" label:"NIK" form:"NIK"`
	NAMA_PEMILIK_USAHA             string         `json:"NAMA_PEMILIK_USAHA" gorm:"column:NAMA_PEMILIK_USAHA;size:50" label:"NAMA_PEMILIK_USAHA" form:"NAMA_PEMILIK_USAHA"`
	TEMPAT_LAHIR_PEMILIK_USAHA     string         `json:"TEMPAT_LAHIR_PEMILIK_USAHA" gorm:"column:TEMPAT_LAHIR_PEMILIK_USAHA;size:100" label:"TEMPAT_LAHIR_PEMILIK_USAHA" form:"TEMPAT_LAHIR_PEMILIK_USAHA"`
	TANGGAL_LAHIR_PEMILIK_USAHA    sql.NullTime   `json:"TANGGAL_LAHIR_PEMILIK_USAHA" gorm:"column:TANGGAL_LAHIR_PEMILIK_USAHA" label:"TANGGAL_LAHIR_PEMILIK_USAHA" form:"TANGGAL_LAHIR_PEMILIK_USAHA"`
	JENIS_KELAMIN_PEMILIK_USAHA    string         `json:"JENIS_KELAMIN_PEMILIK_USAHA" gorm:"column:JENIS_KELAMIN_PEMILIK_USAHA;size:20" label:"JENIS_KELAMIN_PEMILIK_USAHA" form:"JENIS_KELAMIN_PEMILIK_USAHA"`
	PEKERJAAN_PEMILIK_USAHA        string         `json:"PEKERJAAN_PEMILIK_USAHA" gorm:"column:PEKERJAAN_PEMILIK_USAHA;size:200" label:"PEKERJAAN_PEMILIK_USAHA" form:"PEKERJAAN_PEMILIK_USAHA"`
	NOMOR_TELEPON_PEMILIK_USAHA    string         `json:"NOMOR_TELEPON_PEMILIK_USAHA" gorm:"column:NOMOR_TELEPON_PEMILIK_USAHA;size:20" label:"NOMOR_TELEPON_PEMILIK_USAHA" form:"NOMOR_TELEPON_PEMILIK_USAHA"`
	NAMA_PENANGGUNG_JAWAB          string         `json:"NAMA_PENANGGUNG_JAWAB" gorm:"column:NAMA_PENANGGUNG_JAWAB;size:500" label:"NAMA_PENANGGUNG_JAWAB" form:"NAMA_PENANGGUNG_JAWAB"`
	NOMOR_TELEPON_PENANGGUNG_JAWAB string         `json:"NOMOR_TELEPON_PENANGGUNG_JAWAB" gorm:"column:NOMOR_TELEPON_PENANGGUNG_JAWAB;size:20" label:"NOMOR_TELEPON_PENANGGUNG_JAWAB" form:"NOMOR_TELEPON_PENANGGUNG_JAWAB"`
	EMAIL_PENANGGUNG_JAWAB         string         `json:"EMAIL_PENANGGUNG_JAWAB" gorm:"column:EMAIL_PENANGGUNG_JAWAB;size:100" label:"EMAIL_PENANGGUNG_JAWAB" form:"EMAIL_PENANGGUNG_JAWAB"`
	MCC                            string         `json:"MCC" gorm:"column:MCC;size:100" label:"MERCHANT_CATEGORY CODE" form:"MCC"`
	JENIS_USAHA                    string         `json:"JENIS_USAHA" gorm:"column:JENIS_USAHA;size:100" label:"JENIS_USAHA" form:"JENIS_USAHA"`
	PENDAPATAN_USAHA_PERBULAN      string         `json:"PENDAPATAN_USAHA_PERBULAN" gorm:"column:PENDAPATAN_USAHA_PERBULAN;size:100" label:"PENDAPATAN_USAHA_PERBULAN" form:"PENDAPATAN_USAHA_PERBULAN"`
	JENIS_EDC_DIBELI               string         `json:"JENIS_EDC_DIBELI" gorm:"column:JENIS_EDC_DIBELI;size:12" label:"JENIS_EDC_DIBELI"`
	JUMLAH_TRANSAKSI_PERBULAN      string         `json:"JUMLAH_TRANSAKSI_PERBULAN" gorm:"column:JUMLAH_TRANSAKSI_PERBULAN;size:100" label:"JUMLAH_TRANSAKSI_PERBULAN" form:"JUMLAH_TRANSAKSI_PERBULAN"`
	BUSINESS_START_TIME            string         `json:"BUSINESS_START_TIME" gorm:"column:BUSINESS_START_TIME;size:10" label:"BUSINESS_START_TIME" form:"BUSINESS_START_TIME"` // char 821-824
	BUSINESS_END_TIME              string         `json:"BUSINESS_END_TIME" gorm:"column:BUSINESS_END_TIME;size:10" label:"BUSINESS_END_TIME" form:"BUSINESS_END_TIME"`
	FILE_NIB                       string         `json:"FILE_NIB" gorm:"column:FILE_NIB;size:500" label:"FILE_NIB" form:"FILE_NIB"`
	FILE_NPWP                      string         `json:"FILE_NPWP" gorm:"column:FILE_NPWP;size:500" label:"FILE_NPWP" form:"FILE_NPWP"`
	FILE_KTP                       string         `json:"FILE_KTP" gorm:"column:FILE_KTP;size:500" label:"FILE_KTP" form:"FILE_KTP"`
	FILE_SELFIE_KTP                string         `json:"FILE_SELFIE_KTP" gorm:"column:FILE_SELFIE_KTP;size:500" label:"FILE_SELFIE_KTP" form:"FILE_SELFIE_KTP"`
	FILE_FOTO_DEPAN_USAHA          string         `json:"FILE_FOTO_DEPAN_USAHA" gorm:"column:FILE_FOTO_DEPAN_USAHA;size:500" label:"FILE_FOTO_DEPAN_USAHA" form:"FILE_FOTO_DEPAN_USAHA"`
	FILE_SELFIE_LOKASI_USAHA       string         `json:"FILE_SELFIE_LOKASI_USAHA" gorm:"column:FILE_SELFIE_LOKASI_USAHA;size:500" label:"FILE_SELFIE_LOKASI_USAHA" form:"FILE_SELFIE_LOKASI_USAHA"`
	PASSWORD                       string         `json:"PASSWORD" gorm:"column:PASSWORD;size:100" label:"PASSWORD" form:"PASSWORD"`
	PASSWORD_RAW                   string         `json:"PASSWORD_RAW" gorm:"column:PASSWORD_RAW;size:100" label:"MERCHANT_CATEGORY CODE" form:"PASSWORD_RAW"`
	LOGIN_BLOCKED_TIME             time.Time      `json:"LOGIN_BLOCKED_TIME" gorm:"column:LOGIN_BLOCKED_TIME"`
	LOGIN_UNBLOCKED_TIME           time.Time      `json:"LOGIN_UNBLOCKED_TIME" gorm:"column:LOGIN_UNBLOCKED_TIME"`
	LAST_LOGIN                     time.Time      `json:"LAST_LOGIN" gorm:"column:LAST_LOGIN"`
	DIBUAT_PADA                    time.Time      `json:"DIBUAT_PADA" gorm:"column:DIBUAT_PADA;autoCreateTime" form:"DIBUAT_PADA"`
	DISETUJUI_OLEH                 int            `json:"DISETUJUI_OLEH" gorm:"column:DISETUJUI_OLEH;default:0" form:"DISETUJUI_OLEH"`
	DISETUJUI_PADA                 sql.NullTime   `json:"DISETUJUI_PADA" gorm:"column:DISETUJUI_PADA" form:"DISETUJUI_PADA"`
	PENETAPAN_DEVICE_OLEH          int            `json:"PENETAPAN_DEVICE_OLEH" gorm:"column:PENETAPAN_DEVICE_OLEH;default:0" form:"PENETAPAN_DEVICE_OLEH"`
	PENETAPAN_DEVICE_PADA          sql.NullTime   `json:"PENETAPAN_DEVICE_PADA" gorm:"column:PENETAPAN_DEVICE_PADA" form:"PENETAPAN_DEVICE_PADA"`
	DIUBAH_OLEH                    int            `json:"DIUBAH_OLEH" gorm:"column:DIUBAH_OLEH;default:0" form:"DIUBAH_OLEH"`
	DIUBAH_PADA                    time.Time      `json:"DIUBAH_PADA" gorm:"column:DIUBAH_PADA;autoUpdateTime" form:"DIUBAH_PADA"`
	DIHAPUS_PADA                   gorm.DeletedAt `gorm:"column:DIHAPUS_PADA"`
}

func (MerchantApproved) TableName() string {
	return "merchants_approved"
}
func (m *MerchantApproved) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	if m.LOGIN_BLOCKED_TIME.IsZero() {
		m.LOGIN_BLOCKED_TIME = now
	}
	if m.LOGIN_UNBLOCKED_TIME.IsZero() {
		m.LOGIN_UNBLOCKED_TIME = now
	}
	m.LAST_LOGIN = now
	return
}

func (m *MerchantApproved) STATUS_PENDAFTARAN_STRING() string {
	STATUS_LIST := map[int]string{
		-2: "PUTUS KONTRAK",
		-1: "DITOLAK",
		0:  "MENUNGGU PENGISIAN DATA",
		1:  "MENUNGGU PENGECEKAN DATA",
		2:  "MENUNGGU PENETAPAN DEVICE",
		3:  "TERDAFTAR",
	}
	STATUS, ok := STATUS_LIST[m.STATUS_PENDAFTARAN]
	if !ok {
		STATUS = ""
	}
	return STATUS
}
func (MerchantApproved) STATUS_PUTUS_KONTRAK() int {
	return -2
}
func (MerchantApproved) STATUS_DITOLAK() int {
	return -1
}
func (MerchantApproved) STATUS_MENUNGGU_PENGISIAN_DATA() int {
	return 0
}
func (MerchantApproved) STATUS_MENUNGGU_PENGECEKAN_DATA() int {
	return 1
}
func (MerchantApproved) STATUS_MENUNGGU_PENETAPAN_DEVICE() int {
	return 2
}
func (MerchantApproved) STATUS_TERDAFTAR() int {
	return 3
}
func (m *MerchantApproved) BeforeSave(tx *gorm.DB) (err error) {
	// if m.DIBUAT_PADA.IsZero() {
	// 	m.DIBUAT_PADA = time.Now()
	// }
	// if m.DIUBAH_PADA.IsZero() {
	// 	m.DIUBAH_PADA = time.Now()
	// }
	return nil
}
