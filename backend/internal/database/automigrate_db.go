package database

import (
	"fmt"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"

	"gorm.io/gorm"
)

func AutoMigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.UserRole{},
		&model.UserAbilityRule{},
		&model.User{},
	); err != nil {
		return err
	}

	// Pastikan role default tersedia
	db.FirstOrCreate(&model.UserRole{
		ID:    1,
		Title: "Super Admin",
		Name:  "superadmin",
		Icon:  "bx bx-sparkle",
	})

	db.FirstOrCreate(&model.UserRole{
		ID:    2,
		Title: "Default",
		Name:  "default",
		Icon:  "bx bx-radio-circle",
	})

	// Isi ability rule untuk role default
	var count int64
	db.Model(&model.UserAbilityRule{}).Where("role_id IN ?", []int{1, 2}).Count(&count)
	if count == 0 {
		rules := []model.UserAbilityRule{
			{RoleID: 1, Subject: "*", Read: true},
			{RoleID: 2, Subject: "/", Read: true},
			{RoleID: 2, Subject: "/profile", Read: true, Update: true},
		}
		if err := db.Create(&rules).Error; err != nil {
			return fmt.Errorf("failed creating default abilities: %w", err)
		}
	}

	return nil
}
