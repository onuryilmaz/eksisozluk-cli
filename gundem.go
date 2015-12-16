package main

import (
	"github.com/mitchellh/cli"
	"strings"
)

type GundemCommand struct {
	Ui cli.Ui
}

func (c *GundemCommand) Help() string {
	helpText := "Usage: ...."
	return strings.TrimSpace(helpText)
}

func (c *GundemCommand) Run(args []string) int {
	parameter := ParameterFlagHandler(args, c.Ui, c)
	if parameter.Limit == -1 {
		parameter.Limit = 10
	}

	topicList := scraper.GetPopularTopics(parameter)
	WriteTopicList(topicList, parameter)
	return 0
}

func (c *GundemCommand) Synopsis() string {
	return "gundem ile ilgili bilgi"
}
