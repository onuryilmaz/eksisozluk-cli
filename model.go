package main

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
