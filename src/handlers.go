package main

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
	"strings"
)

const EmptyList = "Empty list"

func OnMessageHandler(c tele.Context) error {
	var text = c.Text()

	log.Printf("Message '%s'", text)

	for _, link := range rx.FindAllString(text, -1) {
		storage.Save(link)
	}

	return nil
}

func OnTopicHandler(c tele.Context) error {
	var text = c.Text()
	var payload = c.Message().Payload
	var user = c.Sender()

	log.Printf("Command %s", text)

	count, err := strconv.Atoi(payload)
	if err != nil {
		count = maxLinks
	}
	if count > maxLinks {
		count = maxLinks
	}

	var data []string
	for i, link := range storage.All(count) {
		item := fmt.Sprintf("%d. %s %s", i+1, link.Title, link.Url)
		data = append(data, item)
	}
	message := strings.Join(data, "\n")
	if message == "" {
		message = EmptyList
	}
	_, err = bot.Send(user, message)
	if err != nil {
		return err
	}

	return nil
}

func OnCommandHandler(c tele.Context) error {
	var text = c.Text()

	log.Printf("Command %s", text)

	storage.Clear()

	return nil
}
