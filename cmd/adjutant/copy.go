package main

import (
	"io"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
)

var (
	BufferSize = 2024 * 2024
)

type progressInfo struct {
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

	return func() tea.Msg {
		log.Info(log.Fields{
			"cd-author":  cd.author,
			"tracks":     len(cd.tracks),
			"total-size": totalBytes,
		})

		go func() {
			for i := 0; i < len(cd.tracks); i++ {
				p := progressInfo{
					total:      len(cd.tracks),
					totalBytes: totalBytes,
					doneBytes:  bytesDone,
					done:       i + 1,
					current:    cd.tracks[i].name,
				}
				program.Send(p)
				src := filepath.Join(cfg.source, cd.tracks[i].name)
				dst := filepath.Join(cfg.destination, cd.tracks[i].name)
				copyInternal(src, dst, p, &bytesDone)
			}
			program.Send(completed{
				author:     cd.author,
				title:      cd.title,
				total:      len(cd.tracks),
				totalBytes: totalBytes,
			})
		}()
		return nil
	}
}

func copyInternal(src, dst string, p progressInfo, bytesDone *int64) error {
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
		*bytesDone += int64(n)
		program.Send(progressInfo{
			totalBytes: p.totalBytes,
			doneBytes:  *bytesDone,
			total:      p.total,
			done:       p.done,
			current:    p.current,
		})
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}
