package main

import (
	"fmt"
	"github.com/fatih/color"
)

func WriteEntryList(entryList []Entry) {

	for _, e := range entryList {

		d := color.New(color.FgCyan, color.Bold)
		d.Printf("[%s, %s, %s] ", e.Author, e.Date, e.Id)

		fmt.Println(e.Text)
	}
}

func WriteTopicList(topicList []Topic) {

	for _, t := range topicList {

		fmt.Printf("%s [%d] \n", t.Title, t.Count)
	}
}

func WriteDebeList(debeList []Debe) {

	for _, d := range debeList {

		red := color.New(color.FgRed, color.Bold)
		red.Printf("%s:\n", d.DebeTopic.Title)

		cyan := color.New(color.FgCyan, color.Bold)
		cyan.Printf("[%s, %s, %s] ", d.DebeEntry.Author, d.DebeEntry.Date, d.DebeEntry.Id)

		fmt.Println(d.DebeEntry.Text)
	}
}
