//This is the main package for the "eksisozluk-cli" application
package main

import (
	"github.com/mitchellh/cli"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"strconv"
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

	c := cli.NewCLI("eksisozluk-cli", Version)
	c.Args = os.Args[1:]
	c.Commands = commands
	log.Println("Starting CLI!")
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	log.Println("Closing CLI with status: " + strconv.Itoa(exitStatus))

	os.Exit(exitStatus)
}
