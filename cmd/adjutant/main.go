package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// 'wails dev' should properly launch vite to serve the site
// for live development without needing to seperately launch
// 'npm run dev' or your flavor such as pnpm in the frontend
// directory seperately.

// The comment below chooses what gets packaged with
// the application.

//go:embed all:frontend/build
var assets embed.FS

var app *App

func main() {
	// Create an instance of the app structure
	initConfig()
	app = NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Adjutant",
		Width:            860,
		Height:           500,
		AssetServer:      &assetserver.Options{Assets: assets},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind:             []interface{}{app},
		Debug:            options.Debug{OpenInspectorOnStartup: true},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
