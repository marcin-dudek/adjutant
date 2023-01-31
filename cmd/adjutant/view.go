package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/lipgloss"
	tint "github.com/lrstanley/bubbletint"
	zone "github.com/lrstanley/bubblezone"
)

var (
	theme     = tint.TintGithub
	mainColor = theme.BrightCyan()

	border       = lipgloss.NewStyle().Width(42).BorderForeground(mainColor).Border(lipgloss.DoubleBorder(), false)
	doubleBorder = border.Copy().BorderBottom(true).BorderTop(true).Align(lipgloss.Center).Foreground(mainColor)
	bottomBorder = border.Copy().BorderBottom(true).Foreground(theme.Fg())
	normal       = lipgloss.NewStyle().Foreground(theme.Fg())
	focused      = lipgloss.NewStyle().Foreground(mainColor)
)

func (m model) View() string {
	var b strings.Builder
	fmt.Fprintln(&b, doubleBorder.Render("Adjutant"))

	if m.cd != nil {
		authorStyle := normal
		titleStyle := normal
		blur := func() {
			m.author.Blur()
			m.title.Blur()
		}
		if m.focusIndex == 0 {
			authorStyle = focused
			blur = m.title.Blur
		} else if m.focusIndex == 1 {
			titleStyle = focused
			blur = m.author.Blur
		}
		m.author.TextStyle = authorStyle
		m.title.TextStyle = titleStyle
		blur()

		fmt.Fprintln(&b, zone.Mark("author", m.author.View()))
		fmt.Fprintln(&b, zone.Mark("title", m.title.View()))
		render(&b, "Files  → %d", len(m.cd.tracks))
		render(&b, "Size   → %.2f MB", toMB(m.cd.size))
		renderBottom(&b, "Length → %s", m.cd.length)
	}

	if m.scanning {
		render(&b, "%s reading", m.spinner.View())
	}

	if m.progress != nil {
		fmt.Fprintln(&b, m.progressBar.View())
		render(&b, "Progress → %d/%d", m.progress.done, m.progress.total)
		render(&b, "Progress → %.2f/%.2f MB", toMB(m.progress.doneBytes), toMB(m.progress.totalBytes))
		renderBottom(&b, "Current  → %s", m.progress.current)
	}

	if m.completed != nil {
		fmt.Fprintln(&b, m.progressBar.View())
		render(&b, "Author → %s", m.completed.author)
		render(&b, "Title  → %s", m.completed.title)
		render(&b, "Copied → %d", m.completed.total)
		renderBottom(&b, "Size   → %.2f MB", toMB(m.completed.totalBytes))
	}

	if m.progress == nil && !m.scanning {
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
		fmt.Fprintln(&b,
			scanButton.Render(zone.Mark("scan", "[ SCAN ]")),
			copyButton.Copy().PaddingLeft(8).Render(zone.Mark("copy", "[ COPY ]")),
			exitButton.Copy().PaddingLeft(8).Render(zone.Mark("exit", "[ EXIT ]")))
	}

	return zone.Scan(b.String())
}

func toMB(v int64) float64 {
	return float64(v) / (1000 * 1000)
}

func render(w io.Writer, format string, a ...any) {
	fmt.Fprintln(w, normal.Render(fmt.Sprintf(format, a...)))
}
func renderBottom(w io.Writer, format string, a ...any) {
	fmt.Fprintln(w, bottomBorder.Render(fmt.Sprintf(format, a...)))
}
