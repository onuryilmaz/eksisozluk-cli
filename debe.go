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

	parameter := ParameterFlagHandler(args, c.Ui, c)
	if parameter.Limit == -1 {
		parameter.Limit = 100
	}
	debeList := scraper.GetDEBE(parameter)
	WriteDebeList(debeList, parameter)
	return 0
}

func (c *DebeCommand) Synopsis() string {
	return "debe sonuçları"
}
