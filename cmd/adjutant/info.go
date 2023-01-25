package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	id3 "github.com/bogem/id3v2"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
)

var (
	options id3.Options = id3.Options{Parse: true}
)

func info() tea.Msg {
	go func() {
		var tracks []track
		var size int64
		var total, length time.Duration
		var artist, title string

		filepath.Walk(cfg.source, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				fmt.Println(err)
				return nil
			}
			log.Info(log.Fields{
				"name": info.Name(),
				"size": info.Size(),
				"dir":  info.IsDir(),
				"path": path,
			})

			if !info.IsDir() && filepath.Ext(path) == ".mp3" {
				size += info.Size()
				artist, title, length = mp3details(path)
				total += length
				tracks = append(tracks, track{name: info.Name(), size: info.Size(), path: path})
			}

			return nil
		})

		program.Send(cd{
			author: artist,
			title:  title,
			size:   size,
			tracks: tracks,
			length: total,
		})
	}()

	return scanning{}
}

func mp3details(file string) (string, string, time.Duration) {
	mp3, _ := id3.Open(file, options)
	defer mp3.Close()
	length, _ := strconv.ParseInt(mp3.GetTextFrame(mp3.CommonID("Length")).Text, 10, 32)

	log.Info(log.Fields{
		"length": length,
		"artist": mp3.Artist(),
		"title":  mp3.Title(),
	})

	return mp3.Artist(), mp3.Title(), time.Duration(length) * time.Millisecond
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

type scanning struct {
}
