package main

import (
	"fmt"
	"strings"

	whois "github.com/undiabler/golang-whois"

	"github.com/tomruk/tbauth"
	tb "gopkg.in/tucnak/telebot.v2"
)

func whoisHandler(m *tb.Message) {
	if !tbauth.IsAuthenticated(m.Sender) {
		bot.Send(m.Sender, "You're not authorized!")
		return
	}

	splittedText := strings.Split(m.Text, " ")
	if len(splittedText) != 2 {
		bot.Send(m.Sender, "Usage: /whois args")
		return
	}

	bot.Send(m.Sender, "Who the hell is this? (Whois)")

	result, err := whois.GetWhois(splittedText[1])

	bot.Send(m.Sender, result, tb.NoPreview)

	if err != nil {
		bot.Send(m.Sender, fmt.Sprintf("Also errors occured, details:\n%s", err.Error()))
	}
}
