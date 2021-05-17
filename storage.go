package main

import (
	"bufio"
	"log"
	"os"
)

const filename string = "links.txt"
const write_flags int = os.O_APPEND | os.O_WRONLY | os.O_CREATE
const read_flags int = os.O_RDONLY | os.O_CREATE
const perms = 0600

type Storage struct{}

func (s Storage) getTitle(url string) string {
	text := download(url)
	title := title(text)
	return title
}

func (s Storage) SaveLink(url string) {
	f, err := os.OpenFile(filename, write_flags, perms)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	title := s.getTitle(url)
	item := url + "\n"
	if title != "" {
		item = title + " " + url + "\n"
	}
	if _, err = f.WriteString(item); err != nil {
		panic(err)
	}
	log.Printf("Saved link '%s'", url)
}

func (s Storage) GetLinks(count int) []string {
	var links []string

	f, err := os.OpenFile(filename, read_flags, perms)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan() && i < count; i++ {
		if text := scanner.Text(); text != "" {
			links = append(links, text)
			log.Printf("Read from history '%s'", text)
		}
	}

	return links
}

func (s Storage) Clear() {
	e := os.Truncate(filename, 0)
	log.Printf("Clear '%s'", filename)
	if e != nil {
		log.Fatal(e)
	}
}
