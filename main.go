package main

import (
	"embed"
	"encoding/json"
	"os"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/args"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/logger"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"
)

//go:embed package.json
var embeddedVersion []byte

//go:embed dist/*
var embeddedFiles embed.FS

func main() {
	util.LoadEnv()
	logger.InitLogrus()

	if args.Install() != nil ||
		args.Version(embeddedVersion) != nil {
		return
	}
	var info map[string]any
	if json.Unmarshal(embeddedVersion, &info); info["name"] != "" {
		os.Setenv("APP_NAME", info["name"].(string))
	}
	backend.StartServer(embeddedFiles)
}
