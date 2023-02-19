package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	id3 "github.com/bogem/id3v2"
	tea "github.com/charmbracelet/bubbletea"
	mp3 "github.com/hajimehoshi/go-mp3"
	log "github.com/sirupsen/logrus"
)

var options id3.Options = id3.Options{Parse: true}

func info() tea.Msg {
	go func() {
		var tracks []track
		var size int64
		var total, length time.Duration
		var artist, title string
		log.Info(log.Fields{
			"step": "start-reading-info",
			"path": cfg.source,
		})

		e := filepath.Walk(cfg.source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return nil
			}

			if !info.IsDir() && filepath.Ext(path) == ".mp3" {
				size += info.Size()
				artist, title, length = mp3details(path)
				total += length
				log.Info(log.Fields{
					"step": "reading-files",
					"name": info.Name(), "size": info.Size(), "path": path,
					"length": length, "artist": artist, "title": title,
				})
				tracks = append(tracks, track{name: info.Name(), size: info.Size(), path: path})
			}

			return nil
		})

		log.Info(log.Fields{
			"step": "read-completed", "author": artist, "title": title,
			"size": size, "length": total, "tracks": len(tracks),
		})

		if e != nil {
			log.Error(e)
			program.Send(appError{message: "Not able to read directory"})
		} else {
			program.Send(cd{
				author: artist,
				title:  title,
				size:   size,
				tracks: tracks,
				length: total,
			})
		}
	}()

	return scanning{}
}

func mp3details(file string) (string, string, time.Duration) {
	track, _ := id3.Open(file, options)
	defer track.Close()
	length, e := strconv.ParseInt(track.GetTextFrame(track.CommonID("Length")).Text, 10, 32)
	if e != nil {
		reader, _ := os.Open(file)
		d, _ := mp3.NewDecoder(reader)
		length = (d.Length() / int64((4 * d.SampleRate()))) * 1000
		log.Info(log.Fields{"step": "length-reading", "samples": d.Length(), "rate": d.SampleRate(), "length": length})
	}
	return track.Artist(), track.Title(), time.Duration(length) * time.Millisecond
}

type cd struct {
	author string
	title  string
	tracks []track
	size   int64
	length time.Duration
}

type track struct {
	name string
	path string
	size int64
}

type scanning struct{}

type appError struct {
	message string
}
