package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

var (
	subtle   = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highligh = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	focused     = lipgloss.NewStyle().Foreground(highligh)
	noStyle     = lipgloss.NewStyle()
	borderStyle = lipgloss.NewStyle().Width(40).Border(lipgloss.DoubleBorder(), true, false).BorderForeground(subtle)

	helpStyle    = lipgloss.NewStyle().Foreground(subtle)
	successStyle = lipgloss.NewStyle().Foreground(highligh)

	blurredStyle = noStyle.Copy().PaddingLeft(4).Foreground(subtle)
	focusedStyle = noStyle.Copy().PaddingLeft(4).Foreground(highligh)
	titleStyle   = borderStyle.Copy().PaddingLeft(4)
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
		fmt.Fprintln(&b, helpStyle.Render(fmt.Sprintf("Files  → %d", len(m.cd.tracks))))
		fmt.Fprintln(&b, helpStyle.Render(fmt.Sprintf("Size   → %.2f MB", toMB(m.cd.size))))
		fmt.Fprintln(&b, helpStyle.Render(fmt.Sprintf("Length → %s", m.cd.length)))
	}

	if m.progress != nil {
		fmt.Fprintln(&b, m.progressBar.View())
		fmt.Fprintln(&b, helpStyle.Render(fmt.Sprintf("Progress → %d/%d", m.progress.done, m.progress.total)))
		fmt.Fprintln(&b, helpStyle.Render(fmt.Sprintf("Progress → %.2f/%.2f", toMB(m.progress.doneBytes), toMB(m.progress.totalBytes))))
		fmt.Fprintln(&b, helpStyle.Render(fmt.Sprintf("Current  → %s", m.progress.current)))
	}

	if m.completed != nil {
		fmt.Fprintln(&b, m.progressBar.View())
		fmt.Fprintln(&b, successStyle.Render(fmt.Sprintf("Author → %s", m.completed.author)))
		fmt.Fprintln(&b, successStyle.Render(fmt.Sprintf("Title  → %s", m.completed.title)))
		fmt.Fprintln(&b, successStyle.Render(fmt.Sprintf("Copied → %d", m.completed.total)))
		fmt.Fprintln(&b, successStyle.Render(fmt.Sprintf("Size   → %.2f", toMB(m.completed.totalBytes))))
	}

	if m.progress == nil {
		scanButton := &blurredStyle
		exitButton := &blurredStyle
		if m.focusIndex == 2 {
			scanButton = &focusedStyle
		}
		if m.focusIndex == 3 {
			exitButton = &focusedStyle
		}
		fmt.Fprintf(&b, "\n%s %s\n", scanButton.Render("[ SCAN ]"), exitButton.Render("[ EXIT ]"))
	}

	return b.String()
}

func setFocused(input *textinput.Model) {
	input.PromptStyle = focused
	input.TextStyle = focused
}

func setUnfocused(input *textinput.Model) {
	input.PromptStyle = noStyle
	input.TextStyle = noStyle
	input.Blur()
}

func toMB(v int64) float64 {
	return float64(v) / (1000 * 1000)
}
