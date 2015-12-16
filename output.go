package main

import (
	"fmt"
	"sort"

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
	sortedList := byCount(topicList)
	sort.Sort(sortedList)

	// topics with count higher than mean will be printed in color
	_, idx := sortedList.Mean()
	groupSize := idx / (len(gundemColors) - 1)
	if groupSize == 0 {
		groupSize++
	}

	current := 0

	for i, t := range topicList {
		var prefix, suffix string
		if i <= idx {
			prefix = gundemColors[current]
			suffix = colorEnd

			// color needs to change after first element and after "groupSize" elements
			if current < len(gundemColors)-1 && (i == 0 || i%groupSize == 0) {
				current++
			}
		}

		fmt.Printf("%s%s [%d]%s\n", prefix, t.Title, t.Count, suffix)
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

var gundemColors = []string{
	"\033[38;5;196m",
	"\033[38;5;202m",
	"\033[38;5;208m",
	"\033[38;5;214m",
	"\033[38;5;220m",
}

var colorEnd = "\033[0m"
