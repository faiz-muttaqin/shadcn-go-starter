package main

import (
	"embed"
	"encoding/json"
	"os"

	clerk "github.com/clerk/clerk-sdk-go/v2"
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
	// Set your Clerk Secret Key (from your Dashboard → API Keys)
	clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))

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
