package main

import (

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"golang.org/x/net/html/atom"
	"strconv"

)

var scraper Scraper

type Scraper struct {
	entryMatcher     func(n *html.Node) bool
	authorMatcher    func(n *html.Node) bool
	dateMatcher      func(n *html.Node) bool
	entryListMatcher func(n *html.Node) bool
	topicListMatcher func(n *html.Node) bool
}


func init() {

	entryListMatcher := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "id"), "entry-list")
	}

	entryMatcher := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "class"), "content")
	}

	authorMatcher := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "class"), "entry-author")
	}

	dateMatcher := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "class"), "entry-date")
	}

	topicListMatcher := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "class"), "topic-list")
	}

	scraper = Scraper{entryMatcher, authorMatcher, dateMatcher, entryListMatcher,topicListMatcher}
}

func (s Scraper) findEntries(text string) []entry {
	entryList := make([]entry, 0)

	resp, err := http.Get("https://eksisozluk.com/?q=" + text)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	entryListNode, _ := scrape.Find(root, s.entryListMatcher)

	scrapedEntries := scrape.FindAll(entryListNode, s.entryMatcher)
	for _, scrappedEntry := range scrapedEntries {
		authorNode, authorCheck := scrape.Find(scrappedEntry.Parent, s.authorMatcher)
		dateNode, dateCheck := scrape.Find(scrappedEntry.Parent, s.dateMatcher)

		entry := entry{}
		entry.Text = scrape.Text(scrappedEntry)

		if (authorCheck) {
			entry.Author = scrape.Text(authorNode)
		}
		if (dateCheck) {
			idDate := scrape.Text(dateNode)
			splitted := strings.SplitAfterN(idDate, " ", 2)
			entry.Id = strings.TrimSpace(splitted[0])
			entry.Date = strings.TrimSpace(splitted[1])
		}
		entryList = append(entryList, entry)
	}

	return entryList
}

func (s Scraper) findTopics() []topic {
	topicList := make([]topic, 0)

	resp, err := http.Get("https://eksisozluk.com/basliklar/populer")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	topicListNode, _ := scrape.Find(root, s.topicListMatcher)

	topicLists := scrape.FindAll(topicListNode, scrape.ByTag(atom.Li))
	for _, topicNode := range topicLists {

		topic := topic{}
		titleAndCount := scrape.Text(topicNode)

		countIndex := strings.LastIndex(titleAndCount, " ")

		topic.Title = strings.TrimSpace(titleAndCount[0:countIndex])
		countString := titleAndCount[countIndex:]

		topicCountInt,_ := strconv.Atoi(strings.TrimSpace(countString))
		topic.Count = int64(topicCountInt)

		topicList = append(topicList, topic)
	}

	return topicList
}
