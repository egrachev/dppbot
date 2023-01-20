package main

import (
	"log"
)

type Storage struct{}

func (s Storage) getTitle(url string) string {
	text := download(url)
	title := title(text)
	return title
}

func (s Storage) Save(url string) {
	title := s.getTitle(url)
	db.Create(&Link{Title: title, Url: url})
	log.Printf("Saved link '%s'", url)
}

func (s Storage) All(count int) []Link {
	var links []Link
	db.Find(&links).Limit(count)
	return links
}

func (s Storage) Clear() {
	db.Where("1 = 1").Delete(&Link{})
	log.Printf("Deleted all links")
}
