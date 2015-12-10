package main

import (
	"os"

	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Consul commands.
var commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	commands = map[string]cli.CommandFactory{


		"baslik": func() (cli.Command, error) {
			return &ConfigTestCommand{
				Ui: ui,
			}, nil},

	}
}
