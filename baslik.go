package main

import (
	"github.com/mitchellh/cli"
	"strings"
)

type BaslikCommand struct {
	Ui cli.Ui
}

func (c *BaslikCommand) Help() string {
	helpText := "Usage: ...."
	return strings.TrimSpace(helpText)
}

func (c *BaslikCommand) Run(args []string) int {

	if len(args) < 1 {
		return 1
	}
	baslik := args[0]
	println("Looking for: " + baslik + " ...")

	parameter := ParameterFlagHandler(args[1:], c.Ui, c)

	if parameter.Limit == -1 {
		parameter.Limit = 10
	}
	entryList := scraper.GetEntries(baslik, parameter)

	WriteEntryList(entryList, parameter, baslik)

	return 0
}

func (c *BaslikCommand) Synopsis() string {
	return "baslik adi ile arama"
}
