package main

import (
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
)

var (
	Scan  = 2
	Exit  = 3
	Lines = 3
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
			return m, copyWithArg(*m.cd)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == Exit {
				return m, tea.Quit
			}

			if s == "enter" && m.focusIndex == Scan {
				return m, info
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex = moveDown(m.cd != nil, m.focusIndex)
			} else {
				m.focusIndex = moveUp(m.cd != nil, m.focusIndex)
			}

			var cmd tea.Cmd
			if m.focusIndex == 0 {
				cmd = m.author.Focus()
			} else if m.focusIndex == 1 {
				cmd = m.title.Focus()
			}

			log.Info(log.Fields{"msg": s, "index": m.focusIndex, "step": "after"})
			return m, cmd
		}
	case cd:
		m.cd = &msg
		m.author.SetValue(msg.author)
		m.title.SetValue(msg.title)
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

func moveUp(scanned bool, index int) int {
	if scanned {
		if index == Lines {
			return 0
		}
		return index + 1
	}

	if index == Scan {
		return Exit
	} else {
		return Scan
	}
}

func moveDown(scanned bool, index int) int {
	if scanned {
		if index == 0 {
			return Lines
		}
		return index - 1
	}

	if index == Scan {
		return Exit
	} else {
		return Scan
	}
}
