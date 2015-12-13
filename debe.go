package main

import (
	"github.com/mitchellh/cli"
	"strings"
)

type DebeCommand struct {
	Ui cli.Ui
}

func (c *DebeCommand) Help() string {
	helpText := "Usage: ...."
	return strings.TrimSpace(helpText)
}

func (c *DebeCommand) Run(args []string) int {

	debeList := scraper.GetDEBE()
	WriteDebeList(debeList)
	return 0
}

func (c *DebeCommand) Synopsis() string {
	return "debe sonuçları"
}
