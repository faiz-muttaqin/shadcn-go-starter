package routes

import (
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"

	"github.com/gin-gonic/gin"
)

// MUST BE CALLED BEFORE routes.R = gin.Default()
// OR SETUP GIN
func InitGinMode() {
	if util.IsDevMode() {

	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
