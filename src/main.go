package main

import (
	tele "gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"log"
	"mvdan.cc/xurls/v2"
	"regexp"
)

var db *gorm.DB
var bot *tele.Bot
var rx *regexp.Regexp = xurls.Relaxed()
var maxLinks = 30
var storage = Storage{}

const linksDB = "../bin/links.db"

func main() {
	log.Printf("Start DPP bot")

	initBot()
	initDB()

	bot.Handle(tele.OnText, OnMessageHandler)
	bot.Handle("/topics", OnTopicHandler)
	bot.Handle("/clear", OnCommandHandler)

	bot.Start()

	log.Printf("Stop DPP bot")
}
