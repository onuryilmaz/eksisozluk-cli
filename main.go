package main

import (
	"github.com/mitchellh/cli"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "trace.log",
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})

}

func main() {

	c := cli.NewCLI("eksisozluk-cli", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = commands
	log.Println("Starting CLI!")
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	log.Println("Closing CLI..")

	os.Exit(exitStatus)
}
