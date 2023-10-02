package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Info() CD {
	return info()
}

func (a *App) BookInfo(isbn string) (Book, error) {
	return bookInfo(isbn)
}

func (a *App) Copy(cd CD) {
	copyWithArg(cd)
}

func (a *App) OpenDirectoryDialog(path string) string {
	options := runtime.OpenDialogOptions{DefaultDirectory: path, Title: "Select directory", CanCreateDirectories: true}
	r, _ := runtime.OpenDirectoryDialog(a.ctx, options)
	return r
}

func (a *App) SaveSource(path string) {
	saveSource(path)
}

func (a *App) SaveDestination(path string) {
	saveDestination(path)
}

func (a *App) Configuration() Config {
	return getConfig()
}

func (a *App) emitProgress(data ProgressInfo) {
	runtime.EventsEmit(a.ctx, "copy-progress", data)
}

func (a *App) emitCompleted(data Completed) {
	runtime.EventsEmit(a.ctx, "copy-completed", data)
}
