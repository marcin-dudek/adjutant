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
	borderStyle = lipgloss.NewStyle().Width(40).BorderStyle(lipgloss.NormalBorder()).BorderBottom(true).BorderForeground(subtle)

	helpStyle    = lipgloss.NewStyle().Foreground(subtle)
	blurredStyle = borderStyle.Copy().PaddingLeft(4).Foreground(subtle)
	focusedStyle = borderStyle.Copy().PaddingLeft(4).Foreground(highligh)
	titleStyle   = borderStyle.Copy().MarginTop(1).PaddingLeft(4)
)

func (m model) View() string {
	var b strings.Builder
	fmt.Fprintln(&b, titleStyle.Render("Adjutant"))

	if m.scanned && !m.copying {
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
		fmt.Fprintf(&b, "Files  → %d\n", m.tracks)
		fmt.Fprintf(&b, "Size   → %.2f MB\n", m.sizeInMB)
		fmt.Fprintf(&b, "Length → %s\n", m.length)
		button := &blurredStyle
		if m.focusIndex == Lines {
			button = &focusedStyle
		}
		fmt.Fprintf(&b, "\n%s\n", button.Render("[ EXIT ]"))
	}

	if m.copying {
		fmt.Fprintln(&b, blurredStyle.Render("Progress ..."))
	}

	if m.scanned {
		fmt.Fprintln(&b, helpStyle.Render("alt+s - scan again"))
		fmt.Fprintln(&b, helpStyle.Render("alt+e - copy"))
		fmt.Fprintln(&b, helpStyle.Render("ctrl+c - quit"))
	} else {
		fmt.Fprintln(&b, helpStyle.Render("alt+s - scan"))
		fmt.Fprintln(&b, helpStyle.Render("ctrl+c - quit"))
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
