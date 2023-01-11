package main

import (
	"github.com/charmbracelet/bubbles/progress"
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
			//m.copying = true
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
		m.progress = nil
		m.completed = nil
		return m, nil
	case progressInfo:
		log.Info(log.Fields{
			"progressInfo": msg,
		})
		cmd := m.progressBar.SetPercent(float64(msg.doneBytes) / float64(msg.totalBytes))
		m.progress = &msg
		m.cd = nil
		return m, cmd
	case completed:
		m.focusIndex = Exit
		m.progress = nil
		m.cd = nil
		m.completed = &msg
		return m, nil
	case progress.FrameMsg: // this is for progress bar animation
		progressModel, cmd := m.progressBar.Update(msg)
		m.progressBar = progressModel.(progress.Model)
		return m, cmd
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
