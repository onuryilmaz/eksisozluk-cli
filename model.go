package main

import (
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
	"sort"
)

type Entry struct {
	Text   string
	Author string
	Date   string
	Id     string
}

type Topic struct {
	Title string
	Count int64
	Link  string
}

type Debe struct {
	DebeTopic Topic
	DebeEntry Entry
}
type Parameter struct {
	PageNumber int
	Limit      int
}

type byCount []Topic

func (a byCount) Len() int {
	return len(a)
}

func (a byCount) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byCount) Less(i, j int) bool {
	return a[i].Count > a[j].Count
}

// Mean assumes the list is already sorted.
func (a byCount) Mean() (value int64, index int) {
	for _, t := range a {
		value += t.Count
	}
	n := len(a)
	value = value / int64(n)

	// find the topic that is closest to mean value
	index = sort.Search(n, func(i int) bool { return a[i].Count <= value })
	if index >= n {
		index = n - 1
	}

	return value, index
}

func ParameterFlagHandler(args []string, ui cli.Ui, cli cli.Command) (parameter Parameter) {
	cmdFlags := flag.NewFlagSet("parameter", flag.ContinueOnError)
	cmdFlags.Usage = func() { ui.Output(cli.Help()) }
	var pageNumber, limit int
	cmdFlags.IntVar(&pageNumber, "page", 1, "Sayfa numarasi")
	cmdFlags.IntVar(&limit, "limit", -1, "Limit")
	if err := cmdFlags.Parse(args); err != nil {
		fmt.Println("Error in parameter handling")
	}

	return Parameter{pageNumber, limit}

}
