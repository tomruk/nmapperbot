package main

import (
	"fmt"
	"strings"

	"github.com/tomruk/tbauth"
	tb "gopkg.in/tucnak/telebot.v2"
)

func cloudfailHandler(m *tb.Message) {
	if !tbauth.IsAuthenticated(m.Sender) {
		bot.Send(m.Sender, "You're not authorized!")
		return
	}

	splittedText := strings.Split(m.Text, " ")
	if len(splittedText) <= 1 {
		bot.Send(m.Sender, "Usage: /cloudfail args")
		return
	}

	bot.Send(m.Sender, "Cloudfail scan started, it may take a long time")

	out, err := executeCommand("cloudfail", splittedText, m)

	bot.Send(m.Sender, out, tb.NoPreview)
	if err != nil {
		bot.Send(m.Sender, fmt.Sprintf("Also errors occured, details:\n%s", err.Error()))
	}
}
