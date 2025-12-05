package util

import (
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"

	"gorm.io/gorm"
)

func RemoveEmailSession(db *gorm.DB, email string) {
	updates := map[string]interface{}{
		"last_login": 0,
		"session":    "",
	}
	// Perform the update
	db.Model(&model.Admin{}).Where("email = ?", email).Updates(updates)
}
