package database

import (
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model/modelMerchant"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AutoMigrateDBMerchant(db_merchant *gorm.DB) {

	// Run migrations
	if err := db_merchant.AutoMigrate(
		&modelMerchant.Terminal{},
		&modelMerchant.Merchant{},
		&modelMerchant.MerchantApply{},
		&modelMerchant.TerminatedMerchant{},
		&modelMerchant.LogActivity{},
		&modelMerchant.Feature{},
		&modelMerchant.Role{},
		&modelMerchant.RolePrivilege{},
	); err != nil {
		logrus.Fatal(err)
	}

	// var terminals []modelMerchant.Terminal

	// // Select terminals where email is null
	// db_merchant.Where("email IS NULL").Find(&terminals)

	// for _, terminal := range terminals {
	// 	var merchant modelMerchant.Merchant

	// 	// Find the corresponding merchant
	// 	if err := db_merchant.Where("no = ?", terminal.MerchantId).First(&merchant).Error; err != nil {
	// 		logrus.Println(err)
	// 		continue
	// 	}

	// 	// Update the terminal's email
	// 	newEmail := logrus.Sprintf("%d.%s", terminal.TerminalId, merchant.BusinessContactEmail)
	// 	db_merchant.Model(&terminal).Update("email", newEmail)
	// }

}
