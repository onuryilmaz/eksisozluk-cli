package main

type entry struct {
	Text   string
	Author string
	Date   string
	Id     string
}

type topic struct {
	Title string
	Count int64
}