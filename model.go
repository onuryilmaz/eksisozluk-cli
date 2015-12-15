package main

import "sort"

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
