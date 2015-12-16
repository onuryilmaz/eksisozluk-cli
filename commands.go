package main

import (
	"os"

	"github.com/mitchellh/cli"
)

var commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	commands = map[string]cli.CommandFactory{

		"baslik": func() (cli.Command, error) {
			return &BaslikCommand{
				Ui: ui,
			}, nil
		},

		"gundem": func() (cli.Command, error) {
			return &GundemCommand{
				Ui: ui,
			}, nil
		},

		"debe": func() (cli.Command, error) {
			return &DebeCommand{
				Ui: ui,
			}, nil
		},
	}
}
