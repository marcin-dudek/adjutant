package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bogem/id3v2"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
)

var (
	source  string        = "/home/manek/music/"
	options id3v2.Options = id3v2.Options{Parse: true}
)

func info() tea.Msg {
	files, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	var tracks []track
	var size int64
	var total, length time.Duration
	var artist, title string
	for _, file := range files {
		info, e := file.Info()
		log.Info(log.Fields{
			"name": info.Name(),
			"size": info.Size(),
			"dir":  info.IsDir(),
		})
		if !file.IsDir() && strings.HasSuffix(file.Name(), "mp3") && e == nil {
			size += info.Size()
			artist, title, length = mp3details(file.Name())
			total += length
			tracks = append(tracks, track{name: file.Name(), size: info.Size()})
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

func mp3details(file string) (string, string, time.Duration) {
	mp3, _ := id3v2.Open(filepath.Join(source, file), options)
	defer mp3.Close()
	l, _ := strconv.ParseInt(mp3.GetTextFrame(mp3.CommonID("Length")).Text, 10, 32)

	log.Info(log.Fields{
		"length": l,
		"artist": mp3.Artist(),
		"title":  mp3.Title(),
	})

	return mp3.Artist(), mp3.Title(), time.Duration(l) * time.Millisecond
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
