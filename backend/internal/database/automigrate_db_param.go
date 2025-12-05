package database

import (
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model/modelParam"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AutoMigrateDBParam(db_param *gorm.DB) {

	// Run migrations
	if err := db_param.AutoMigrate(
		&modelParam.ResponseCodeReversal{},
		&modelParam.ResponseCodeTrx{},
		&modelParam.SMTP{},
	); err != nil {
		logrus.Fatal(err)
	}

	var smtpCount int64
	db_param.Model(&modelParam.SMTP{}).Count(&smtpCount)
	if smtpCount == 0 {
		var smtps = []modelParam.SMTP{
			{
				HOST:     "smtp.gmail.com",
				PORT:     587,
				EMAIL:    "service4@csna4u.com",
				PASSWORD: "service.12345",
				SENDER:   "smartwebindonesia@csna.com",
			},
		}

		// Perform batch insert
		db_param.Create(&smtps)

		for _, smtp := range smtps {
			// Access IDs after insert
			logrus.Println("Insert New smtp  with ID : ", smtp.ID)
		}
	}
}
