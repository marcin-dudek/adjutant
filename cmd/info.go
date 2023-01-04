package main

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bogem/id3v2"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
)

var (
	path string = "/home/manek/music/"
)

func info() tea.Msg {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	options := id3v2.Options{Parse: true}

	var tracks []track
	var size int64
	var total time.Duration
	var artist, title string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), "mp3") {
			size += file.Size()
			mp3, err := id3v2.Open(filepath.Join(path, file.Name()), options)
			if err != nil {
				return err
			}
			defer mp3.Close()
			l, _ := strconv.ParseInt(mp3.GetTextFrame(mp3.CommonID("Length")).Text, 10, 32)

			log.Info(log.Fields{
				"length": l,
				"artist": mp3.Artist(),
				"title":  mp3.Title(),
			})
			artist = mp3.Artist()
			title = mp3.Title()

			total += time.Duration(l) * time.Millisecond
			tracks = append(tracks, track{name: file.Name(), size: file.Size()})
		}
	}

	return cd{
		author: artist,
		title:  title,
		size:   size,
		tracks: tracks,
		length: total,
	}
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
	size int64
}
