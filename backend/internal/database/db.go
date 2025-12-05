package database

import (
	"os"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DB_TRX *gorm.DB
var DB_PARAM *gorm.DB
var DB_MERCHANT *gorm.DB
var DB_ODOO *gorm.DB
var DB_ODOO_CS *gorm.DB

func Init() error {
	var err error
	if DB, err = util.ConnectToSQLDB(
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		util.Getenv("DB_PORT", "0"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
	); err != nil {
		logrus.Printf("Database setup failed: %v", err)
		return err
	}
	go func() {
		if err := AutoMigrateDB(DB); err != nil {
			logrus.Fatalf("Auto migrate database failed: %v", err)
		}
	}()
	return nil
}
