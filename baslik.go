package main


import (
	"strings"
	"github.com/mitchellh/cli"

)

// ConfigTestCommand is a Command implementation that is used to
// verify config files
type ConfigTestCommand struct {
	Ui cli.Ui
}

func (c *ConfigTestCommand) Help() string {
	helpText := "Usage: ...."
	return strings.TrimSpace(helpText)
}

func (c *ConfigTestCommand) Run(args []string) int {

	if (len(args) < 1) {
		return 1
	}
	baslik := args[0]
	println("Looking for: " + baslik + "...")
	sanitizedText := strings.Replace(baslik, " ", "%20%", -1)

	entryList := scraper.findEntries(sanitizedText)
	WriteEntryList(entryList)

	return 0
}

func (c *ConfigTestCommand) Synopsis() string {
	return "baslik adi ile arama"
}
