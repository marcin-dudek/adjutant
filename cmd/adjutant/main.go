package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
	log "github.com/sirupsen/logrus"
)

func (m model) Init() tea.Cmd {
	return nil
}

var (
	program *tea.Program
	cfg     config
)

func main() {
	f, err := initLogFile()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("AppStarted")
	cfg = initConfig()
	zone.NewGlobal()
	program = tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := program.Run(); err != nil {
		log.Error(err)
		fmt.Printf("could not start program: %s\n", err)
		return
	}

	log.Info("AppExited")
}

func initLogFile() (*os.File, error) {
	dir := "logs"
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	file := fmt.Sprintf("%s/log_%s.txt", dir, time.Now().Format("20060102"))
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}
