package main
import (
	"fmt"
	"github.com/fatih/color"
)
func WriteEntryList(entryList []entry) {

	for _, e := range entryList {

		d := color.New(color.FgCyan, color.Bold)
		d.Printf("[%s, %s, %s] ", e.Author, e.Date, e.Id)

		fmt.Println(e.Text)
	}
}
