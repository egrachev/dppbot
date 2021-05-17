package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"mvdan.cc/xurls/v2"

	tb "gopkg.in/tucnak/telebot.v2"
)

var max_links int = 30
var storage Storage = Storage{}

func main() {
	var token = os.Getenv("TOKEN")
	log.Printf("Start DPP bot")

	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		log.Printf("Message '%s'", m.Text)

		rx := xurls.Relaxed()
		for _, link := range rx.FindAllString(m.Text, -1) {
			storage.SaveLink(link)
		}
	})

	b.Handle("/topics", func(m *tb.Message) {
		log.Printf("Command %s", m.Text)

		count, err := strconv.Atoi(m.Payload)
		if err != nil {
			count = max_links
		}
		if count > max_links {
			count = max_links
		}

		var data []string
		for i, link := range storage.GetLinks(count) {
			item := fmt.Sprintf("%d. %s", i+1, link)
			data = append(data, item)
		}
		message := strings.Join(data, "\n")
		b.Send(m.Chat, message)
	})

	b.Handle("/clear", func(m *tb.Message) {
		log.Printf("Command %s", m.Text)

		storage.Clear()
	})

	b.Start()
}
