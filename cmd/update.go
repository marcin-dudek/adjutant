package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

var (
	Lines = 2
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "alt+s":
			return m, info
		case "alt+e":
			m.copying = true
			return m, nil

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == Lines {
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > Lines {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = Lines
			}

			var cmd tea.Cmd
			if m.focusIndex == 0 {
				cmd = m.author.Focus()
			} else if m.focusIndex == 1 {
				cmd = m.title.Focus()
			}

			return m, cmd
		}
	case cd:
		m.scanned = true
		m.author.SetValue(msg.author)
		m.title.SetValue(msg.title)
		m.tracks = len(msg.tracks)
		m.scanned = true
		m.sizeInMB = float64(msg.size) / (1024 * 1024)
		m.length = msg.length
		return m, nil
	}

	// Handle character input and blinking
	var cmd tea.Cmd
	if m.focusIndex == 0 {
		m.author, cmd = m.author.Update(msg)
	} else if m.focusIndex == 1 {
		m.title, cmd = m.title.Update(msg)
	}

	return m, cmd
}
