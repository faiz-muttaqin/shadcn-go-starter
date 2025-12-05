package database

import (
	"encoding/json"
	"os"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AutoMigrateDBDummy(db *gorm.DB) {
	db.AutoMigrate(&model.ExamplePerson{})

	// Define the file path
	filePath := "./web/assets/json/table-datatable.json"

	// Read the file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Println("Error :" + err.Error())
	}

	// Unmarshal the JSON data into a map
	var jsonData map[string]interface{}
	if err := json.Unmarshal(fileContent, &jsonData); err != nil {
		logrus.Println("Error :" + err.Error())
	}

	// Extract data and convert to structs
	var merchants []model.ExamplePerson
	if data, ok := jsonData["data"].([]interface{}); ok {
		for _, item := range data {
			if itemMap, ok := item.(map[string]interface{}); ok {
				var merchant model.ExamplePerson
				merchant.ID = uint(itemMap["id"].(float64)) // JSON numbers are float64
				merchant.Avatar = itemMap["avatar"].(string)
				merchant.FullName = itemMap["full_name"].(string)
				merchant.Post = itemMap["post"].(string)
				merchant.Email = itemMap["email"].(string)
				merchant.City = itemMap["city"].(string)
				merchant.StartDate = itemMap["start_date"].(string)
				merchant.Salary = itemMap["salary"].(string)
				merchant.Age = itemMap["age"].(string)
				merchant.Experience = itemMap["experience"].(string)
				merchant.Status = int(itemMap["status"].(float64))

				merchants = append(merchants, merchant)
			}
		}
	}

	// Check if data exists
	var merchantCount int64
	db.Model(&model.ExamplePerson{}).Count(&merchantCount)
	if merchantCount == 0 {
		// Perform batch insert
		if err := db.Create(&merchants).Error; err != nil {
			logrus.Println("Error :" + err.Error())
		}

		for _, merchant := range merchants {
			// Access IDs after insert
			logrus.Println("Inserted Merchant with ID : ", merchant.ID)
		}
	}

}
