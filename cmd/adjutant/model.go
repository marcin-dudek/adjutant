package main

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
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

	return model{
		author:      author,
		title:       title,
		progressBar: progress.New(progress.WithDefaultGradient()),
		progress:    nil,
		focusIndex:  2,
		cd:          nil,
		completed:   nil,
	}
}
