package modelMerchant

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type MerchantApply struct {
	ID                                         int64          `json:"id" gorm:"column:id;primarykey"`
	MID                                        int64          `json:"MID" gorm:"column:MID"`
	STATUS_PENDAFTARAN                         int            `json:"STATUS_PENDAFTARAN" gorm:"column:STATUS_PENDAFTARAN;default:0"`
	KETERANGAN_PENDAFTARAN                     string         `json:"KETERANGAN_PENDAFTARAN" gorm:"column:KETERANGAN_PENDAFTARAN;size:1000"`
	NAMA_AGREGATOR                             string         `json:"NAMA_AGREGATOR" gorm:"column:NAMA_AGREGATOR;size:200"`
	NOMOR_PENDAFTARAN                          string         `json:"NOMOR_PENDAFTARAN" gorm:"column:NOMOR_PENDAFTARAN;size:500" label:"NOMOR_PENDAFTARAN"`
	EMAIL                                      string         `json:"EMAIL" gorm:"column:EMAIL;size:500" label:"EMAIL"`
	SETUJU_KETENTUAN_LAYANAN                   bool           `json:"SETUJU_KETENTUAN_LAYANAN" gorm:"column:SETUJU_KETENTUAN_LAYANAN" label:"SETUJU_KETENTUAN_LAYANAN"`
	SETUJU_PENGGUNAAN_INFORMASI                bool           `json:"SETUJU_PENGGUNAAN_INFORMASI" gorm:"column:SETUJU_PENGGUNAAN_INFORMASI" label:"SETUJU_PENGGUNAAN_INFORMASI"`
	JENIS_PENDAFTARAN                          string         `json:"JENIS_PENDAFTARAN" gorm:"column:JENIS_PENDAFTARAN;size:100" label:"JENIS_PENDAFTARAN"`
	BENTUK_BADAN_HUKUM                         string         `json:"BENTUK_BADAN_HUKUM" gorm:"column:BENTUK_BADAN_HUKUM;size:500" label:"BENTUK_BADAN_HUKUM"`
	NAMA_PEMEGANG_SAHAM                        string         `json:"NAMA_PEMEGANG_SAHAM" gorm:"column:NAMA_PEMEGANG_SAHAM;size:100" label:"NAMA_PEMEGANG_SAHAM"`
	NIB                                        string         `json:"NIB" gorm:"column:NIB;size:100" label:"NIB"`
	NAMA_USAHA                                 string         `json:"NAMA_USAHA" gorm:"column:NAMA_USAHA;size:500" label:"NAMA_USAHA"`
	ALAMAT_USAHA                               string         `json:"ALAMAT_USAHA" gorm:"column:ALAMAT_USAHA;size:500" label:"ALAMAT_USAHA"`
	KODE_POS_FULL                              string         `json:"KODE_POS_FULL" gorm:"column:KODE_POS_FULL;size:500" label:"KODE_POS_FULL"`
	KODE_POS                                   string         `json:"KODE_POS" gorm:"column:KODE_POS;size:20" label:"KODE_POS"`
	PROVINCE                                   string         `json:"PROVINCE" gorm:"column:PROVINCE;size:100" label:"PROVINSI"`
	CITY                                       string         `json:"CITY" gorm:"column:CITY;size:100" label:"KOTA/KABUPATEN"`
	DISTRICT                                   string         `json:"DISTRICT" gorm:"column:DISTRICT;size:100" label:"KECAMATAN"`
	SUBDISTRICT                                string         `json:"SUBDISTRICT" gorm:"column:SUBDISTRICT;size:100" label:"DESA/KELURAHAN"`
	NOMOR_TELEPON_USAHA                        string         `json:"NOMOR_TELEPON_USAHA" gorm:"column:NOMOR_TELEPON_USAHA;size:20" label:"NOMOR_TELEPON_USAHA"`
	NAMA_PEMILIK_REKENING                      string         `json:"NAMA_PEMILIK_REKENING" gorm:"column:NAMA_PEMILIK_REKENING;size:50" label:"NAMA_PEMILIK_REKENING"`
	NOMOR_REKENING                             string         `json:"NOMOR_REKENING" gorm:"column:NOMOR_REKENING;size:50" label:"NOMOR_REKENING"`
	BANK_CODE                                  string         `json:"BANK_CODE" gorm:"column:BANK_CODE;size:50" label:"BANK_CODE"`
	BANK_PENERBIT_REKENING                     string         `json:"BANK_PENERBIT_REKENING" gorm:"column:BANK_PENERBIT_REKENING;size:50" label:"BANK_PENERBIT_REKENING"`
	NIK                                        string         `json:"NIK" gorm:"column:NIK;size:50" label:"NIK"`
	NAMA_PEMILIK_USAHA                         string         `json:"NAMA_PEMILIK_USAHA" gorm:"column:NAMA_PEMILIK_USAHA;size:50" label:"NAMA_PEMILIK_USAHA"`
	NAMA_PEMILIK_USAHA_ALIAS                   string         `json:"NAMA_PEMILIK_USAHA_ALIAS" gorm:"column:NAMA_PEMILIK_USAHA_ALIAS;size:50" label:"NAMA_PEMILIK_USAHA_ALIAS"`
	TEMPAT_LAHIR_PEMILIK_USAHA                 string         `json:"TEMPAT_LAHIR_PEMILIK_USAHA" gorm:"column:TEMPAT_LAHIR_PEMILIK_USAHA;size:100" label:"TEMPAT_LAHIR_PEMILIK_USAHA"`
	TANGGAL_LAHIR_PEMILIK_USAHA                sql.NullTime   `json:"TANGGAL_LAHIR_PEMILIK_USAHA" gorm:"column:TANGGAL_LAHIR_PEMILIK_USAHA" label:"TANGGAL_LAHIR_PEMILIK_USAHA"`
	JENIS_KELAMIN_PEMILIK_USAHA                string         `json:"JENIS_KELAMIN_PEMILIK_USAHA" gorm:"column:JENIS_KELAMIN_PEMILIK_USAHA;size:20" label:"JENIS_KELAMIN_PEMILIK_USAHA"`
	PEKERJAAN_PEMILIK_USAHA                    string         `json:"PEKERJAAN_PEMILIK_USAHA" gorm:"column:PEKERJAAN_PEMILIK_USAHA;size:200" label:"PEKERJAAN_PEMILIK_USAHA"`
	NOMOR_TELEPON_PEMILIK_USAHA                string         `json:"NOMOR_TELEPON_PEMILIK_USAHA" gorm:"column:NOMOR_TELEPON_PEMILIK_USAHA;size:20" label:"NOMOR_TELEPON_PEMILIK_USAHA"`
	NAMA_PENANGGUNG_JAWAB                      string         `json:"NAMA_PENANGGUNG_JAWAB" gorm:"column:NAMA_PENANGGUNG_JAWAB;size:500" label:"NAMA_PENANGGUNG_JAWAB"`
	NOMOR_TELEPON_PENANGGUNG_JAWAB             string         `json:"NOMOR_TELEPON_PENANGGUNG_JAWAB" gorm:"column:NOMOR_TELEPON_PENANGGUNG_JAWAB;size:20" label:"NOMOR_TELEPON_PENANGGUNG_JAWAB"`
	EMAIL_PENANGGUNG_JAWAB                     string         `json:"EMAIL_PENANGGUNG_JAWAB" gorm:"column:EMAIL_PENANGGUNG_JAWAB;size:100" label:"EMAIL_PENANGGUNG_JAWAB"`
	MCC                                        string         `json:"MCC" gorm:"column:MCC;size:100" label:"MERCHANT_CATEGORY CODE"`
	JENIS_USAHA                                string         `json:"JENIS_USAHA" gorm:"column:JENIS_USAHA;size:100" label:"JENIS_USAHA"`
	PENDAPATAN_USAHA_PERBULAN                  string         `json:"PENDAPATAN_USAHA_PERBULAN" gorm:"column:PENDAPATAN_USAHA_PERBULAN;size:100" label:"PENDAPATAN_USAHA_PERBULAN"`
	JUMLAH_TRANSAKSI_PERBULAN                  string         `json:"JUMLAH_TRANSAKSI_PERBULAN" gorm:"column:JUMLAH_TRANSAKSI_PERBULAN;size:100" label:"JUMLAH_TRANSAKSI_PERBULAN"`
	JUMLAH_PEMBELIAN_EDC                       int            `json:"JUMLAH_PEMBELIAN_EDC" gorm:"column:JUMLAH_PEMBELIAN_EDC" label:"JUMLAH_PEMBELIAN_EDC"`
	JENIS_EDC_DIBELI                           string         `json:"JENIS_EDC_DIBELI" gorm:"column:JENIS_EDC_DIBELI;size:12" label:"JENIS_EDC_DIBELI"`
	FILE_NIB                                   string         `json:"FILE_NIB" gorm:"column:FILE_NIB;size:500" label:"FILE_NIB"`
	FILE_AKTE_PENDIRIAN_DAN_PENGESAHAN         string         `json:"FILE_AKTE_PENDIRIAN_DAN_PENGESAHAN" gorm:"column:FILE_AKTE_PENDIRIAN_DAN_PENGESAHAN;size:1000" label:"FILE_AKTE_PENDIRIAN_DAN_PENGESAHAN"`
	FILE_AKTE_PERUBAHAN_TERKINI_DAN_PENGESAHAN string         `json:"FILE_AKTE_PERUBAHAN_TERKINI_DAN_PENGESAHAN" gorm:"column:FILE_AKTE_PERUBAHAN_TERKINI_DAN_PENGESAHAN;size:1000" label:"FILE_AKTE_PERUBAHAN_TERKINI_DAN_PENGESAHAN"`
	FILE_SURAT_KUASA_PIC                       string         `json:"FILE_SURAT_KUASA_PIC" gorm:"column:FILE_SURAT_KUASA_PIC;size:1000" label:"FILE_AKTE_PERUBAHAN_TERKINI_DAN_PENGESAHAN"`
	FILE_NPWP                                  string         `json:"FILE_NPWP" gorm:"column:FILE_NPWP;size:500" label:"FILE_NPWP"`
	FILE_KTP                                   string         `json:"FILE_KTP" gorm:"column:FILE_KTP;size:500" label:"FILE_KTP"`
	FILE_SELFIE_KTP                            string         `json:"FILE_SELFIE_KTP" gorm:"column:FILE_SELFIE_KTP;size:500" label:"FILE_SELFIE_KTP"`
	FILE_FOTO_DEPAN_USAHA                      string         `json:"FILE_FOTO_DEPAN_USAHA" gorm:"column:FILE_FOTO_DEPAN_USAHA;size:500" label:"FILE_FOTO_DEPAN_USAHA"`
	FILE_SELFIE_LOKASI_USAHA                   string         `json:"FILE_SELFIE_LOKASI_USAHA" gorm:"column:FILE_SELFIE_LOKASI_USAHA;size:500" label:"FILE_SELFIE_LOKASI_USAHA"`
	PASSWORD                                   string         `json:"PASSWORD" gorm:"column:PASSWORD;size:500" label:"PASSWORD"`
	PASSWORD_RAW                               string         `json:"PASSWORD_RAW" gorm:"column:PASSWORD_RAW;size:100" label:"MERCHANT_CATEGORY CODE"`
	DIBUAT_PADA                                time.Time      `json:"DIBUAT_PADA" gorm:"column:DIBUAT_PADA;autoCreateTime"`
	DISETUJUI_OLEH                             int            `json:"DISETUJUI_OLEH" gorm:"column:DISETUJUI_OLEH;default:0"`
	DISETUJUI_PADA                             sql.NullTime   `json:"DISETUJUI_PADA" gorm:"column:DISETUJUI_PADA"`
	DITOLAK_OLEH                               int            `json:"DITOLAK_OLEH" gorm:"column:DITOLAK_OLEH;default:0"`
	DITOLAK_PADA                               sql.NullTime   `json:"DITOLAK_PADA" gorm:"column:DITOLAK_PADA"`
	PENETAPAN_DEVICE_OLEH                      int            `json:"PENETAPAN_DEVICE_OLEH" gorm:"column:PENETAPAN_DEVICE_OLEH;default:0"`
	PENETAPAN_DEVICE_PADA                      sql.NullTime   `json:"PENETAPAN_DEVICE_PADA" gorm:"column:PENETAPAN_DEVICE_PADA"`
	DIUBAH_OLEH                                int            `json:"DIUBAH_OLEH" gorm:"column:DIUBAH_OLEH;default:0"`
	DIUBAH_PADA                                time.Time      `json:"DIUBAH_PADA" gorm:"column:DIUBAH_PADA;autoUpdateTime"`
	DIHAPUS_PADA                               gorm.DeletedAt `gorm:"column:DIHAPUS_PADA"`
}

func (MerchantApply) TableName() string {
	return "merchants_apply"
}

func (m *MerchantApply) STATUS_PENDAFTARAN_STRING() string {
	STATUS_LIST := map[int]string{
		-2: "DITUNDA",
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
func (MerchantApply) STATUS_DITOLAK() int {
	return -1
}
func (MerchantApply) STATUS_MENUNGGU_PENGISIAN_DATA() int {
	return 0
}
func (MerchantApply) STATUS_MENUNGGU_PENGECEKAN_DATA() int {
	return 1
}
func (MerchantApply) STATUS_MENUNGGU_PENETAPAN_DEVICE() int {
	return 2
}
func (MerchantApply) STATUS_DITUNDA() int {
	return -2
}
func (MerchantApply) STATUS_TERDAFTAR() int {
	return 3
}
func (m *MerchantApply) BeforeCreate(tx *gorm.DB) (err error) {
	if m.STATUS_PENDAFTARAN == 0 {
		m.STATUS_PENDAFTARAN = m.STATUS_MENUNGGU_PENGISIAN_DATA()
	}
	if m.DIBUAT_PADA.IsZero() {
		m.DIBUAT_PADA = time.Now()
	}
	if m.DIUBAH_PADA.IsZero() {
		m.DIUBAH_PADA = time.Now()
	}
	return
}
func (m *MerchantApply) BeforeSave(tx *gorm.DB) (err error) {
	// if m.DIBUAT_PADA.IsZero() {
	// 	m.DIBUAT_PADA = time.Now()
	// }
	// if m.DIUBAH_PADA.IsZero() {
	// 	m.DIUBAH_PADA = time.Now()
	// }
	return nil
}
