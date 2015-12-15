package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {

	c := cli.NewCLI("app", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = commands

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(exitStatus)
}
