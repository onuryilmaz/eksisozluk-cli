package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jfrazelle/go/canonical/json"
	"github.com/mattn/go-colorable"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

var gundemColors = []string{
	"\033[38;5;196m",
	"\033[38;5;202m",
	"\033[38;5;208m",
	"\033[38;5;214m",
	"\033[38;5;220m",
}

var colorEnd = "\033[0m"

// WriteEntryList handles writing entries based on parameters
func WriteEntryList(entryList []Entry, parameter Parameter, baslik string) {

	if parameter.Output == "console" {
		writeEntryListConsole(entryList)
	} else if parameter.Output == "json" {
		writeJSON(entryList, parameter, "debe")
	} else {
		log.Println("No supported output!")
	}
}

// WriteDebeList handles writing "Debe"s based on parameters
func WriteDebeList(debeList []Debe, parameter Parameter) {

	if parameter.Output == "console" {
		writeDebeListConsole(debeList)
	} else if parameter.Output == "json" {
		writeJSON(debeList, parameter, "debe")
	} else {
		log.Println("No supported output!")
	}
}

// WriteTopicList handles writing topics based on parameters
func WriteTopicList(topicList []Topic, parameter Parameter) {
	if parameter.Output == "console" {
		writeTopicListConsole(topicList)
	} else if parameter.Output == "json" {
		writeJSON(topicList, parameter, "gundem")
	} else {
		log.Println("No supported output!")
	}
}

func fileNameHandler(baslik string, parameter Parameter) string {

	fileName := baslik

	fileName = fileName + "_limit_" + strconv.Itoa(parameter.Limit)

	if parameter.PageNumber > 1 {
		fileName = fileName + "_page_" + strconv.Itoa(parameter.PageNumber)
	}
	if parameter.Sukela {
		fileName = fileName + "_sukela"
	}

	return fileName
}

func writeJSON(data interface{}, parameter Parameter, baslik string) {

	j, jerr := json.MarshalIndent(data, "", "  ")
	if jerr != nil {
		fmt.Println("jerr:", jerr.Error())
	}

	fileName := fileNameHandler(baslik, parameter)

	ioutil.WriteFile(fileName+".json", j, 0777)
	log.Println("Writing to file: " + fileName + ".json")

}

func writeEntryListConsole(entryList []Entry) {

	for _, e := range entryList {

		d := color.New(color.FgCyan, color.Bold)
		d.Printf("[%s, %s, %s] ", e.Author, e.Date, e.ID)

		fmt.Println(e.Text)
	}
}


func writeTopicListConsole(topicList []Topic) {
	sortedList := byCount(topicList)
	sort.Sort(sortedList)

	// topics with count higher than mean will be printed in color
	_, idx := sortedList.Mean()
	groupSize := idx / (len(gundemColors) - 1)
	if groupSize == 0 {
		groupSize++
	}

	current := 0
	writer := colorable.NewColorableStdout()

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

		fmt.Fprintf(writer, "%s%s [%d]%s\n", prefix, t.Title, t.Count, suffix)
	}
}


func writeDebeListConsole(debeList []Debe) {

	for _, d := range debeList {

		red := color.New(color.FgRed, color.Bold)
		red.Printf("%s:\n", d.DebeTopic.Title)

		cyan := color.New(color.FgCyan, color.Bold)
		cyan.Printf("[%s, %s, %s] ", d.DebeEntry.Author, d.DebeEntry.Date, d.DebeEntry.ID)

		fmt.Println(d.DebeEntry.Text)
	}
}

