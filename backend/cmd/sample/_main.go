package main

import (
	"fmt"
	"os"
	"path/filepath"

	"embed"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/database"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/middleware"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/routes"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/args"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/clr"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/docs"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/logger"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/version"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:embed docs.json
var docsFile []byte

//go:embed package.json
var embeddedVersion []byte

//go:embed dist/*
var embeddedFiles embed.FS

const (
	DOCS_URLPATH  = "/docs"
	DOCS_FILENAME = "docs.json"
)

func main() {
	isDevMode := util.IsDevMode()
	if isDevMode {
		fmt.Println(clr.BgYellow("Development Mode"))
	}
	util.LoadEnv()
	logger.InitLogrus()

	if args.Install() != nil ||
		args.Version(embeddedVersion) != nil {
		return
	}

	database.Init()

	//HANDLE LOG WRITING
	ginLogFile, err := logger.InitGinLogger()
	if err != nil {
		logrus.Fatal(err)
	}
	defer ginLogFile.Close()

	// routes.R.Use(middleware.CacheControlMiddleware())
	routes.InitGinMode()
	routes.R = gin.Default()
	routes.R.Use(logger.GinLoggerMiddleware(ginLogFile))
	routes.R.Use(middleware.SecurityControlMiddleware())
	routes.R.Use(cors.Default())

	routes.WebSocketRoutes()
	routes.Routes()

	r1 := routes.EmbedFilesHandler(embeddedFiles, isDevMode)
	r2 := docs.ServeSwaggerDocs(routes.R, DOCS_URLPATH, DOCS_FILENAME, docsFile)

	logrus.Info(clr.BgCyan(" http://localhost" + util.Getenv("APP_LOCAL_HOST", ":18080") + " "))

	if isDevMode { // Only run when in Go Run Mode
		go version.Generate(filepath.Join(filepath.Dir(filepath.Clean(os.Getenv("APP_DIR"))), "package.json"))
		go docs.GenerateSwaggerDoc(routes.R, DOCS_URLPATH, DOCS_FILENAME, append(r1, r2...)...)
	}
	if err := routes.R.Run(os.Getenv("APP_LOCAL_HOST")); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
