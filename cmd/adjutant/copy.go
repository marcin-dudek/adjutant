package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
)

var BufferSize = 2024 * 2024

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
	time       time.Duration
}

func copyWithArg(cd cd, author, title string) tea.Cmd {
	var totalBytes, bytesDone int64
	for _, t := range cd.tracks {
		totalBytes += t.size
	}
	t := time.Now()
	return func() tea.Msg {
		log.Info(log.Fields{
			"step":       "copy-progress",
			"author":     author,
			"title":      title,
			"tracks":     len(cd.tracks),
			"total-size": totalBytes,
		})

		destination := filepath.Join(cfg.destination, fmt.Sprintf("%s - %s", author, title))
		if _, err := os.Stat(destination); errors.Is(err, os.ErrNotExist) {
			if e := os.Mkdir(destination, os.ModePerm); e != nil {
				log.Error(e)
				return nil
				// return appError{message: "Failed to create directory. '" + destination + "'."}
			}
		}

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
				dst := filepath.Join(destination, cd.tracks[i].name)
				copyInternal(cd.tracks[i].path, dst, p, &bytesDone)
			}

			program.Send(completed{
				author:     cd.author,
				title:      cd.title,
				total:      len(cd.tracks),
				totalBytes: totalBytes,
				time:       time.Since(t).Truncate(10 * time.Millisecond),
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
	}
	return nil
}
