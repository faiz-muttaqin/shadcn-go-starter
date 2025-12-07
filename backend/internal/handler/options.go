package handler

import (
	"net/http"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/database"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Option struct {
	Label string `json:"label"`
	Value any    `json:"value"`
	Icon  string `json:"icon,omitempty"`
}

func GetOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := c.Query("data")
		options := []Option{}
		title := cases.Title(language.English).String(data)
		switch data {
		case "status":
			options = []Option{
				{Label: "Active", Value: "active"},
				{Label: "Inactive", Value: "inactive"},
				{Label: "Invited", Value: "invited"},
				{Label: "Suspended", Value: "suspended"},
			}

		case "role":
			var roles []model.UserRole
			database.DB.Find(&roles)
			for _, r := range roles {
				options = append(options, Option{
					Label: cases.Title(language.English).String(r.Title),
					Value: r.Name,
					Icon:  r.Icon,
				})
			}
		default:
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "options data not found",
				"error":   "options data not found",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    options,
			"title":   title,
		})
	}
}
