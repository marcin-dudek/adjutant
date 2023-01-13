package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/lipgloss"
	tint "github.com/lrstanley/bubbletint"
)

var (
	theme = tint.TintGithub

	border  = lipgloss.NewStyle().Width(40).Border(lipgloss.DoubleBorder(), true, false).BorderForeground(theme.BrightCyan())
	normal  = lipgloss.NewStyle().Foreground(theme.Fg())
	focused = lipgloss.NewStyle().Foreground(theme.BrightCyan())

	titleStyle = border.Copy().PaddingLeft(4).Foreground(theme.BrightCyan())
)

func (m model) View() string {
	var b strings.Builder
	fmt.Fprintln(&b, titleStyle.Render("Adjutant"))

	if m.cd != nil {
		authorStyle := normal
		titleStyle := normal
		cmd := func() {
			m.author.Blur()
			m.title.Blur()
		}
		if m.focusIndex == 0 {
			authorStyle = focused
			cmd = m.title.Blur
		} else if m.focusIndex == 1 {
			titleStyle = focused
			cmd = m.author.Blur
		}
		m.author.TextStyle = authorStyle
		m.title.TextStyle = titleStyle
		cmd()

		fmt.Fprintln(&b, m.author.View())
		fmt.Fprintln(&b, m.title.View())
		render(&b, "Files  → %d", len(m.cd.tracks))
		render(&b, "Size   → %.2f MB", toMB(m.cd.size))
		render(&b, "Length → %s", m.cd.length)
	}

	if m.progress != nil {
		fmt.Fprintln(&b, m.progressBar.View())
		render(&b, "Progress → %d/%d", m.progress.done, m.progress.total)
		render(&b, "Progress → %.2f/%.2f MB", toMB(m.progress.doneBytes), toMB(m.progress.totalBytes))
		render(&b, "Current  → %s", m.progress.current)
	}

	if m.completed != nil {
		fmt.Fprintln(&b, m.progressBar.View())
		render(&b, "Author → %s", m.completed.author)
		render(&b, "Title  → %s", m.completed.title)
		render(&b, "Copied → %d", m.completed.total)
		render(&b, "Size   → %.2f MB", toMB(m.completed.totalBytes))
	}

	if m.progress == nil {
		scanButton := &normal
		copyButton := &normal
		exitButton := &normal
		if m.focusIndex == ScanIndex {
			scanButton = &focused
		}
		if m.focusIndex == CopyIndex {
			copyButton = &focused
		}
		if m.focusIndex == ExitIndex {
			exitButton = &focused
		}
		fmt.Fprintf(&b, "%s %s %s\n", scanButton.Render("[ SCAN ]"), copyButton.Render("[ COPY ]"), exitButton.Render("[ EXIT ]"))
	}

	return b.String()
}

func toMB(v int64) float64 {
	return float64(v) / (1000 * 1000)
}

func render(w io.Writer, format string, a ...any) {
	fmt.Fprintln(w, normal.Render(fmt.Sprintf(format, a...)))
}
