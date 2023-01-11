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
	var author = textinput.New()
	author.Prompt = "Author → "
	author.Placeholder = "Author"
	author.CharLimit = 64
	author.SetCursorMode(textinput.CursorStatic)
	author.Focus()

	var title = textinput.New()
	title.Prompt = "Title  → "
	title.Placeholder = "Title"
	title.CharLimit = 64
	title.SetCursorMode(textinput.CursorStatic)

	progressBar := progress.New(
		progress.WithColorProfile(termenv.TrueColor),
		progress.WithScaledGradient(tint.Hex(theme.Fg()), tint.Hex(theme.BrightCyan())),
	)

	return model{
		author:      author,
		title:       title,
		progressBar: progressBar,
		progress:    nil,
		focusIndex:  2,
		cd:          nil,
		completed:   nil,
	}
}
