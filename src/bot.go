package main

import (
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func initBot() {
	var err error

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err = tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
}
