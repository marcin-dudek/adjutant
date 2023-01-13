package main

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tint "github.com/lrstanley/bubbletint"
	"github.com/muesli/termenv"
)

type model struct {
	focusIndex  int
	author      textinput.Model
	title       textinput.Model
	progressBar progress.Model
	progress    *progressInfo
	cd          *cd
	completed   *completed
}

func initialModel() model {
	progressBar := progress.New(
		progress.WithColorProfile(termenv.TrueColor),
		progress.WithScaledGradient(tint.Hex(theme.Fg()), tint.Hex(theme.BrightCyan())),
	)

	return model{
		author:      newInput("Author → "),
		title:       newInput("Title  → "),
		progressBar: progressBar,
		progress:    nil,
		focusIndex:  2,
		cd:          nil,
		completed:   nil,
	}
}

func newInput(prompt string) textinput.Model {
	input := textinput.New()
	input.Prompt = prompt
	input.CharLimit = 64
	input.PromptStyle = normal
	input.SetCursorMode(textinput.CursorStatic)
	return input
}
