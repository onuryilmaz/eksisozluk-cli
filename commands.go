package main

import (
	"os"

	"github.com/mitchellh/cli"
)

// commands holds the command implementations
var commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	commands = map[string]cli.CommandFactory{

		"baslik": func() (cli.Command, error) {
			return &BaslikCommand{
				cli: EksiSozlukCLICommand{ui},
			}, nil
		},

		"gundem": func() (cli.Command, error) {
			return &GundemCommand{
				cli: EksiSozlukCLICommand{ui},
			}, nil
		},

		"debe": func() (cli.Command, error) {
			return &DebeCommand{
				cli: EksiSozlukCLICommand{ui},
			}, nil
		},
	}
}
