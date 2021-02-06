package main

import (
	"fmt"
	"strings"

	"github.com/tomruk/tbauth"
	tb "gopkg.in/tucnak/telebot.v2"
)

func digHandler(m *tb.Message) {
	if !tbauth.IsAuthenticated(m.Sender) {
		bot.Send(m.Sender, "You're not authorized!")
		return
	}

	splittedText := strings.Split(m.Text, " ")
	if len(splittedText) <= 1 {
		bot.Send(m.Sender, "Usage: /dig args")
		return
	}

	bot.Send(m.Sender, "Digging")

	out, err := executeCommand("dig", splittedText, m)

	bot.Send(m.Sender, out, tb.NoPreview)
	if err != nil {
		bot.Send(m.Sender, fmt.Sprintf("Also errors occured, details:\n%s", err.Error()))
	}
}
