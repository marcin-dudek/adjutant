package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

var BufferSize = 2024 * 2024

type ProgressInfo struct {
	Total      int
	Done       int
	TotalBytes int64
	DoneBytes  int64
	Current    string
}

type Completed struct {
	Author     string
	Title      string
	Total      int
	TotalBytes int64
	Time       time.Duration
	Path       string
}

func copyWithArg(cd CD) {
	var totalBytes, bytesDone int64 = 0, 0
	for _, t := range cd.Tracks {
		totalBytes += t.Size
	}
	t := time.Now()
	log.Info(log.Fields{
		"step":       "copy-progress",
		"author":     cd.Author,
		"title":      cd.Title,
		"tracks":     len(cd.Tracks),
		"total-size": totalBytes,
	})

	destination := filepath.Join(cfg.Destination, fmt.Sprintf("%s - %s", cd.Author, cd.Title))
	if _, err := os.Stat(destination); errors.Is(err, os.ErrNotExist) {
		if e := os.Mkdir(destination, os.ModePerm); e != nil {
			log.Error(e)
			//return appError{message: "Failed to create directory. '" + destination + "'."}
		}
	}

	go func() {
		for i := 0; i < len(cd.Tracks); i++ {
			p := ProgressInfo{
				Total:      len(cd.Tracks),
				TotalBytes: totalBytes,
				DoneBytes:  bytesDone,
				Done:       i + 1,
				Current:    cd.Tracks[i].Name,
			}
			log.Info(p)
			app.emitProgress(p)

			dst := filepath.Join(destination, cd.Tracks[i].Name)
			copyInternal(cd.Tracks[i].Path, dst, p, &bytesDone)
		}

		app.emitCompleted(Completed{
			Author:     cd.Author,
			Title:      cd.Title,
			Total:      len(cd.Tracks),
			TotalBytes: totalBytes,
			Time:       time.Since(t).Truncate(10 * time.Millisecond),
			Path:       destination,
		})
	}()
}

func copyInternal(src, dst string, p ProgressInfo, bytesDone *int64) error {
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
		time.Sleep(700 * time.Millisecond)

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
		*bytesDone += int64(n)
		app.emitProgress(ProgressInfo{
			TotalBytes: p.TotalBytes,
			DoneBytes:  *bytesDone,
			Total:      p.Total,
			Done:       p.Done,
			Current:    p.Current,
		})
	}
	return nil
}
