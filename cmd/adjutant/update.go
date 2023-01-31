package main

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
	log "github.com/sirupsen/logrus"
)

var (
	ScanIndex = 2
	CopyIndex = 3
	ExitIndex = 4
	Lines     = 4
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
			return m, copyWithArg(*m.cd, m.author.Value(), m.title.Value())

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == ExitIndex {
				return m, tea.Quit
			}

			if s == "enter" && m.focusIndex == CopyIndex {
				return m, copyWithArg(*m.cd, m.author.Value(), m.title.Value())
			}

			if s == "enter" && m.focusIndex == ScanIndex {
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

			log.Info(log.Fields{"step": "update", "msg": s, "index": m.focusIndex})
			return m, cmd

		case "left":
			if m.focusIndex > 1 { // navigate only when on buttons
				m.focusIndex = moveDown(m.cd != nil, m.focusIndex)
				return m, nil
			}
		case "right":
			if m.focusIndex > 1 { // navigate only when on buttons
				m.focusIndex = moveUp(m.cd != nil, m.focusIndex)
				return m, nil
			}
		}
	case scanning:
		m.scanning = true
		m.cd = nil
		m.progress = nil
		m.completed = nil
		return m, m.spinner.Tick
	case cd:
		m.cd = &msg
		m.author.SetValue(msg.author)
		m.title.SetValue(msg.title)
		m.progress = nil
		m.completed = nil
		m.scanning = false
		return m, nil
	case progressInfo:
		log.Info(log.Fields{
			"step":         "progress-info",
			"progressInfo": msg,
		})
		cmd := m.progressBar.SetPercent(float64(msg.doneBytes) / float64(msg.totalBytes))
		m.progress = &msg
		m.cd = nil
		return m, cmd
	case completed:
		m.focusIndex = ExitIndex
		m.progress = nil
		m.cd = nil
		m.completed = &msg
		return m, nil
	case progress.FrameMsg: // this is for progress bar animation
		progressModel, cmd := m.progressBar.Update(msg)
		m.progressBar = progressModel.(progress.Model)
		return m, cmd
	case spinner.TickMsg: // this is for spinner animation
		spinner, cmd := m.spinner.Update(msg)
		m.spinner = spinner
		return m, cmd
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return m, nil
		}
		var cmd tea.Cmd
		if zone.Get("scan").InBounds(msg) {
			m.focusIndex = ScanIndex
			cmd = info
		} else if zone.Get("exit").InBounds(msg) {
			cmd = tea.Quit
		} else if zone.Get("copy").InBounds(msg) {
			if m.cd != nil {
				m.focusIndex = CopyIndex
				cmd = copyWithArg(*m.cd, m.author.Value(), m.title.Value())
			}
		} else if zone.Get("author").InBounds(msg) {
			m.focusIndex = 0
			cmd = m.author.Focus()
		} else if zone.Get("title").InBounds(msg) {
			m.focusIndex = 1
			cmd = m.title.Focus()
		}
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

	if index == ScanIndex {
		return ExitIndex
	} else {
		return ScanIndex
	}
}

func moveDown(scanned bool, index int) int {
	if scanned {
		if index == 0 {
			return Lines
		}
		return index - 1
	}

	if index == ScanIndex {
		return ExitIndex
	} else {
		return ScanIndex
	}
}
