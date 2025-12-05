package database

import (
	"gorm.io/gorm"
)

func AutoMigrateDBACPS(acps_db *gorm.DB) {

	// Run migrations
	// if err := acps_db.AutoMigrate(
	// 	&pg_acps.MDB_MERCHANT_INFO{},
	// ); err != nil {
	// 	logrus.Fatal(err)
	// }

}
