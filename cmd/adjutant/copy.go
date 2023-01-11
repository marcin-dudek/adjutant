package main

import (
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	BufferSize = 1_000_000
)

type progress struct {
	total      int
	done       int
	totalBytes int64
	doneBytes  int64
	current    string
}

type completed struct {
	author     string
	title      string
	total      int
	totalBytes int64
}

func copyWithArg(cd cd) tea.Cmd {
	var totalBytes, bytesDone int64
	for _, t := range cd.tracks {
		totalBytes += t.size
	}
	p := progress{
		total:      len(cd.tracks),
		done:       0,
		totalBytes: totalBytes,
		doneBytes:  0,
		current:    cd.tracks[0].name,
	}

	return func() tea.Msg {
		log.Info(log.Fields{
			"cd-author": cd.author,
			"tracks":    len(cd.tracks),
		})
		go func() {
			for i := 0; i < len(cd.tracks); i++ {
				p := progress{
					total:      p.total,
					totalBytes: p.totalBytes,
					doneBytes:  bytesDone,
					done:       i + 1,
					current:    cd.tracks[i].name,
				}
				program.Send(p)
				copyInternal("/home/manek/music/"+cd.tracks[i].name, "/home/manek/src/tmp/"+cd.tracks[i].name)
				bytesDone += cd.tracks[i].size

				program.Send(progress{
					total:      p.total,
					totalBytes: p.totalBytes,
					doneBytes:  bytesDone,
					done:       i + 1,
					current:    cd.tracks[i].name,
				})
			}
			program.Send(completed{
				author:     cd.author,
				title:      cd.title,
				total:      len(cd.tracks),
				totalBytes: totalBytes,
			})
		}()
		program.Send(p)
		return nil
	}
}

func copyInternal(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		os.Remove(dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	buf := make([]byte, BufferSize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}
