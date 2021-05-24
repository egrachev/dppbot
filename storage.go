package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const links_filename string = "links.txt"
const write_flags int = os.O_APPEND | os.O_WRONLY | os.O_CREATE
const read_flags int = os.O_RDONLY | os.O_CREATE
const links_perms = 0600

type Storage struct{}

func (s Storage) getTitle(url string) string {
	text := download(url)
	title := title(text)
	return title
}

func (s Storage) SaveLink(url string) {
	f, err := os.OpenFile(links_filename, write_flags, links_perms)
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

	f, err := os.OpenFile(links_filename, read_flags, links_perms)
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

func (s Storage) Backup() {
	input, err := ioutil.ReadFile(links_filename)
	if err != nil {
		fmt.Println("Error read links file", links_filename)
		fmt.Println(err)
	}

	now := time.Now()
	links_backup := links_filename + "." + now.Format("20060102150405")
	err = ioutil.WriteFile(links_backup, input, links_perms)
	fmt.Println("Create backup", links_backup)
	if err != nil {
		fmt.Println("Error creating", links_backup)
		fmt.Println(err)
	}
}

func (s Storage) Clear() {
	s.Backup()
	e := os.Truncate(links_filename, 0)
	log.Printf("Clear '%s'", links_filename)
	if e != nil {
		log.Fatal(e)
	}
}
