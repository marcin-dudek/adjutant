package main

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	focusIndex int
	author     textinput.Model
	title      textinput.Model
	tracks     int
	sizeInMB   float64
	scanned    bool
	copying    bool
	progress   int
	length     time.Duration
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
		author:     author,
		title:      title,
		scanned:    false,
		copying:    false,
		focusIndex: 2,
	}
}
