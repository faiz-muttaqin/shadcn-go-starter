package backend

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"embed"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/database"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/middleware"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/routes"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/clr"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/docs"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/kvstore"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/logger"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/version"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:embed docs.json
var docsFile []byte

const (
	DOCS_URLPATH  = "/docs"
	DOCS_FILENAME = "docs.json"
)

func StartServer(embeddedFiles embed.FS) {
	isDevMode := util.IsDevMode()
	if isDevMode {
		fmt.Println(clr.BgYellow(clr.TextBlack("Development Mode")))
	}
	database.Init()
	go func() {
		kvstore.RDB = kvstore.InitRedis(
			os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
			os.Getenv("REDIS_PASSWORD"),
			os.Getenv("REDIS_DB"),
		)
	}()
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
	routes.R.Use(middleware.Security())
	routes.R.Use(cors.Default())

	routes.WebSocketRoutes()
	routes.Routes()

	r1 := routes.NoRouteDefaultFiles(embeddedFiles, isDevMode)
	r2 := docs.ServeSwaggerDocs(routes.R, DOCS_URLPATH, DOCS_FILENAME, docsFile)

	logrus.Info(clr.TextCyan(" http://localhost" + util.Getenv("APP_LOCAL_HOST", ":8173") + " "))

	if isDevMode { // Only run when in Go Run Mode
		go version.Generate(filepath.Join(os.Getenv("APP_DIR"), "package.json"))
		go docs.GenerateSwaggerDoc(routes.R, filepath.Join(util.ThisFileDir(runtime.Caller(0)), DOCS_FILENAME), append(r1, r2...)...)
	}
	if err := routes.R.Run(os.Getenv("APP_LOCAL_HOST")); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
