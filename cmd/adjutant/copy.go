package main

import (
	log "github.com/sirupsen/logrus"

	tea "github.com/charmbracelet/bubbletea"
)

type progress struct {
	total   int
	done    int
	current string
}

func copyWithArg(cd cd) tea.Cmd {
	p := progress{
		total:   len(cd.tracks),
		done:    0,
		current: cd.tracks[0].name,
	}

	return func() tea.Msg {
		log.Info(log.Fields{
			"cd-author": cd.author,
			"tracks":    len(cd.tracks),
		})
		program.Send(p)
		return nil
	}
}
