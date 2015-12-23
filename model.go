package main

import (
	"errors"
	"flag"
	"sort"

	"github.com/mitchellh/cli"
)

type EksiSozlukCLICommand struct {
	UI cli.Ui
}

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
	Sukela     bool
	Output     string
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
	n := len(a)
	if n < 1 {
		return
	}

	for _, t := range a {
		value += t.Count
	}
	value = value / int64(n)

	// find the topic that is closest to mean value
	index = sort.Search(n, func(i int) bool { return a[i].Count <= value })
	if index >= n {
		index = n - 1
	}

	return
}

func ParameterFlagHandler(args []string, eksiCLI interface{}, c EksiSozlukCLICommand) (parameter Parameter, err error) {
	cmdFlags := flag.NewFlagSet("parameter", flag.ContinueOnError)
	var pageNumber, limit int
	var sukelaMod bool
	var output string
	cmdFlags.IntVar(&pageNumber, "page", 1, "Sayfa numarasi")
	cmdFlags.IntVar(&limit, "limit", -1, "Limit")
	cmdFlags.BoolVar(&sukelaMod, "sukela", false, "Sukela mod")
	cmdFlags.StringVar(&output, "output", "console", "Output mode")
	cmdFlags.Usage = func() { c.UI.Output(eksiCLI.(cli.Command).Help()) }

	if err := cmdFlags.Parse(args); err != nil {
		return parameter, errors.New("Error in parameter handling")
	}
	return Parameter{pageNumber, limit, sukelaMod, output}, nil

}
