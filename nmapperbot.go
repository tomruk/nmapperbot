package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/mbndr/figlet4go"
	"github.com/tomruk/tbauth"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	bot *tb.Bot
	err error
)

func main() {
	ascii := figlet4go.NewAsciiRender()
	rendered, _ := ascii.Render("Nmapperbot")

	fmt.Println(aurora.Bold(aurora.Blue(rendered)))

	token := os.Getenv("NMAPPERBOT_TOKEN")
	passphrase := os.Getenv("NMAPPERBOT_PASSPHRASE")

	if token == "" {
		fmt.Println("NMAPPERBOT_TOKEN should not be empty")
		os.Exit(1)
	}

	if passphrase == "" {
		fmt.Println("NMAPPERBOT_PASSPHRASE should not be empty")
		os.Exit(1)
	}

	tbauth.Passphrase = &passphrase

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		fmt.Print(aurora.Bold(aurora.Red("\r-- Exiting")))
		os.Exit(0)
	}()

	bot, err = tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, "Welcome to the Nmapper")
	})

	bot.Handle("/auth", func(m *tb.Message) {
		splittedText := strings.Split(m.Text, " ")
		if len(splittedText) != 2 {
			bot.Send(m.Sender, "Usage: /auth passphrase")
			return
		}

		resp := tbauth.Authenticate(m.Sender, splittedText[1])
		switch resp {
		case 0:
			bot.Send(m.Sender, "Successfully authorized 👍🏻")
		case 1:
			bot.Send(m.Sender, "You're already authorized 😁")
		case 2:
			bot.Send(m.Sender, "I didn't like your passphrase bro  👎🏻")
		}
	})

	// Network related
	bot.Handle("/nmap", nmapHandler)

	// DNS and domain related
	bot.Handle("/whois", whoisHandler)
	bot.Handle("/dig", digHandler)
	bot.Handle("/nslookup", nslookupHandler)

	// Automatic tools
	bot.Handle("/cloudfail", cloudfailHandler)

	fmt.Println(aurora.Bold(aurora.Green("-> Bot started")))
	bot.Start()
}
