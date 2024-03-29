package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	id3 "github.com/bogem/id3v2"
	mp3 "github.com/hajimehoshi/go-mp3"
	log "github.com/sirupsen/logrus"
)

func info() CD {
	var tracks []Track
	var size int64
	var total, length time.Duration
	var artist, title string
	log.Info(log.Fields{
		"step": "start-reading-info",
		"path": cfg.Source,
	})

	filepath.Walk(cfg.Source, func(path string, info os.FileInfo, err error) error {
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
			tracks = append(tracks, Track{Name: info.Name(), Size: info.Size(), Path: path})
		}

		return nil
	})

	log.Info(log.Fields{
		"step": "read-completed", "author": artist, "title": title,
		"size": size, "length": total, "tracks": len(tracks),
	})

	return CD{
		Author: artist,
		Title:  title,
		Size:   size,
		Tracks: tracks,
		Length: total,
	}
}

func mp3details(file string) (string, string, time.Duration) {
	track, _ := id3.Open(file, id3.Options{Parse: true})
	defer track.Close()
	length, e1 := strconv.ParseInt(track.GetTextFrame(track.CommonID("Length")).Text, 10, 32)
	if e1 != nil {
		reader, _ := os.Open(file)
		d, e2 := mp3.NewDecoder(reader)
		if e2 != nil {
			length = (d.Length() / int64((4 * d.SampleRate()))) * 1000
			log.Info(log.Fields{"step": "length-reading", "samples": d.Length(), "rate": d.SampleRate(), "length": length})
		} else {
			log.Warn(log.Fields{"step": "length-reading", "error": e2})
		}
	}
	return track.Artist(), track.Title(), time.Duration(length) * time.Millisecond
}

type CD struct {
	Author string
	Title  string
	Tracks []Track
	Size   int64
	Length time.Duration
}

type Track struct {
	Name string
	Path string
	Size int64
}

type appError struct {
	message string
}
