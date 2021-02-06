package main

import (
	"os/exec"

	tb "gopkg.in/tucnak/telebot.v2"
)

func executeCommand(cmd string, args []string, m *tb.Message) (string, error) {
	args = append(args[:0], args[1:]...)

	out, err := exec.Command(cmd, args...).CombinedOutput()
	return string(out), err
}
