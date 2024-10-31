package main

import (
	"SuperSiteLogsApp/handlers"
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed frontend/dist/*
var assets embed.FS

func main() {
	appInstance := &App{}
	logsInstance := &handlers.Logs{}
	homeInstance := &handlers.Home{}
	logUtilInstance := &handlers.LogsUtils{}
	keyPathInstance := &handlers.KeyPathSelect{}
	syncLogsInstance := &handlers.SyncLogs{}
	logViewInstance := &handlers.LogView{}

	err := wails.Run(&options.App{
		Title:         "Hello World",
		Width:         800,
		Height:        600,
		MinWidth:      800,
		MinHeight:     600,
		Frameless:     false,
		DisableResize: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			appInstance,
			logsInstance,
			homeInstance,
			logUtilInstance,
			keyPathInstance,
			syncLogsInstance,
			logViewInstance,
		},
		OnStartup: appInstance.Startup,
	})

	if err != nil {
		log.Fatal(err)
	}
}
