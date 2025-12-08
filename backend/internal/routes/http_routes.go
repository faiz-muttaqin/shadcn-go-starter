package routes

import (
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/database"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/handler"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func Routes() {
	// Endpoint login API
	r := R.Group(util.GetPathOnly(util.Getenv("VITE_BACKEND", "/api")))

	r.GET("/options", handler.GetOptions())
	// r.GET("/users/hehe", handler.GET_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"Role"}))
	api := r.Group("")
	{
		api.OPTIONS("/auth/login", handler.GetAuthLogin())
		api.GET("/auth/login", handler.GetAuthLogin())
		api.GET("/auth/logout", handler.GetAuthLogout())
		api.GET("/auth/verify", handler.VerifyAuth()) // Test auth endpoint
		api.GET("/roles", handler.GetRoles())
		api.GET("/users", handler.GET_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
		api.POST("/users", handler.POST_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
		api.PATCH("/users", handler.PATCH_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
		api.PUT("/users", handler.PUT_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
		api.DELETE("/users", handler.DELETE_DEFAULT_TableDataHandler(database.DB, &model.User{}))

		// Theme endpoints
		api.GET("/themes", handler.GET_DEFAULT_TableDataHandler(database.DB, &model.Theme{}, []string{}))
		api.POST("/themes", handler.POST_DEFAULT_TableDataHandler(database.DB, &model.Theme{}, []string{}))
		api.PATCH("/themes", handler.PATCH_DEFAULT_TableDataHandler(database.DB, &model.Theme{}, []string{}))
		api.DELETE("/themes", handler.DELETE_DEFAULT_TableDataHandler(database.DB, &model.Theme{}))
	}
}
