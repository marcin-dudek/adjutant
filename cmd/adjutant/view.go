package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
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
		if m.focusIndex == 0 {
			setFocused(&m.author)
			setUnfocused(&m.title)
		} else if m.focusIndex == 1 {
			setFocused(&m.title)
			setUnfocused(&m.author)
		} else {
			setUnfocused(&m.author)
			setUnfocused(&m.title)
		}

		fmt.Fprintln(&b, m.author.View())
		fmt.Fprintln(&b, m.title.View())
		fmt.Fprintln(&b, normal.Render(fmt.Sprintf("Files  → %d", len(m.cd.tracks))))
		fmt.Fprintln(&b, normal.Render(fmt.Sprintf("Size   → %.2f MB", toMB(m.cd.size))))
		fmt.Fprintln(&b, normal.Render(fmt.Sprintf("Length → %s", m.cd.length)))
	}

	if m.progress != nil {
		fmt.Fprintln(&b, m.progressBar.View())
		fmt.Fprintln(&b, normal.Render(fmt.Sprintf("Progress → %d/%d", m.progress.done, m.progress.total)))
		fmt.Fprintln(&b, normal.Render(fmt.Sprintf("Progress → %.2f/%.2f MB", toMB(m.progress.doneBytes), toMB(m.progress.totalBytes))))
		fmt.Fprintln(&b, normal.Render(fmt.Sprintf("Current  → %s", m.progress.current)))
	}

	if m.completed != nil {
		fmt.Fprintln(&b, m.progressBar.View())
		fmt.Fprintln(&b, normal.Render(fmt.Sprintf("Author → %s", m.completed.author)))
		fmt.Fprintln(&b, normal.Render(fmt.Sprintf("Title  → %s", m.completed.title)))
		fmt.Fprintln(&b, normal.Render(fmt.Sprintf("Copied → %d", m.completed.total)))
		fmt.Fprintln(&b, normal.Render(fmt.Sprintf("Size   → %.2f MB", toMB(m.completed.totalBytes))))
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

func setFocused(input *textinput.Model) {
	input.PromptStyle = normal
	input.TextStyle = focused
}

func setUnfocused(input *textinput.Model) {
	input.PromptStyle = normal
	input.TextStyle = normal
	input.Blur()
}

func toMB(v int64) float64 {
	return float64(v) / (1000 * 1000)
}
