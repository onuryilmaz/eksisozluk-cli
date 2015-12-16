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

	topicList := scraper.GetPopularTopics(parameter)
	WriteTopicList(topicList)
	return 0
}

func (c *GundemCommand) Synopsis() string {
	return "gundem ile ilgili bilgi"
}
