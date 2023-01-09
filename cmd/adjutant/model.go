package main

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	focusIndex int
	author     textinput.Model
	title      textinput.Model
	copying    bool
	progress   int
	cd         *cd
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
		copying:    false,
		focusIndex: 2,
		cd:         nil,
	}
}
