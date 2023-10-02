package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Author struct {
	//URL  string `json:"url"`
	Name string `json:"name"`
}

type Cover struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type Publisher struct {
	Name string `json:"name"`
}

type Book struct {
	URL         string      `json:"url"`
	Key         string      `json:"key"`
	Title       string      `json:"title"`
	Authors     []Author    `json:"authors"`
	Publishers  []Publisher `json:"publishers"`
	PublishDate string      `json:"publish_date"`
	Cover       Cover       `json:"cover"`
}

func bookInfo(isbn string) (Book, error) {
	url := fmt.Sprintf("http://openlibrary.org/api/books?bibkeys=ISBN:%s&jscmd=data&format=json", isbn)
	r, e := http.Get(url)
	if e != nil {
		return Book{}, e
	}
	m := make(map[string]Book)
	json.NewDecoder(r.Body).Decode(&m)

	return m["ISBN:"+isbn], nil
}
