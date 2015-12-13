package main

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var scraper Scraper

type Scraper struct {
	entryMatcher     func(n *html.Node) bool
	authorMatcher    func(n *html.Node) bool
	dateMatcher      func(n *html.Node) bool
	entryListMatcher func(n *html.Node) bool
	topicListMatcher func(n *html.Node) bool
	indexListMatcher func(n *html.Node) bool
	contentMatcher   func(n *html.Node) bool
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

	indexListMatcher := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "id"), "index-section")
	}

	contentMatcher := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "id"), "content-body")
	}

	scraper = Scraper{entryMatcher, authorMatcher, dateMatcher, entryListMatcher, topicListMatcher, indexListMatcher, contentMatcher}
}

func (s Scraper) GetEntries(text string) []Entry {
	return s.getEntries("https://eksisozluk.com/?q=" + url.QueryEscape(text))
}
func (s Scraper) getEntries(eksiURL string) []Entry {

	resp, err := http.Get(eksiURL)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	entryList := make([]Entry, 0)

	entryListNode, found := scrape.Find(root, s.entryListMatcher)

	if found == false {
		return entryList
	}

	scrapedEntries := scrape.FindAll(entryListNode, s.entryMatcher)

	if len(scrapedEntries) == 0 {
		return entryList
	}
	for _, scrappedEntry := range scrapedEntries {
		authorNode, authorCheck := scrape.Find(scrappedEntry.Parent, s.authorMatcher)
		dateNode, dateCheck := scrape.Find(scrappedEntry.Parent, s.dateMatcher)

		entry := Entry{}
		entry.Text = scrape.Text(scrappedEntry)

		if authorCheck {
			entry.Author = scrape.Text(authorNode)
		}
		if dateCheck {
			idDate := scrape.Text(dateNode)
			splitted := strings.SplitAfterN(idDate, " ", 2)
			entry.Id = strings.TrimSpace(splitted[0])
			entry.Date = strings.TrimSpace(splitted[1])
		}
		entryList = append(entryList, entry)
	}

	return entryList
}

func (s Scraper) GetPopularTopics() []Topic {

	resp, err := http.Get("https://eksisozluk.com/basliklar/populer")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	topicListNode, _ := scrape.Find(root, s.indexListMatcher)

	return s.getTopics(topicListNode)

}

func (s Scraper) getTopics(node *html.Node) []Topic {
	topicList := make([]Topic, 0)

	topicListNode, _ := scrape.Find(node, s.topicListMatcher)

	topicLists := scrape.FindAll(topicListNode, scrape.ByTag(atom.Li))
	for _, topicNode := range topicLists {

		topic := Topic{}
		topicLink, _ := scrape.Find(topicNode, scrape.ByTag(atom.A))
		topic.Link = "https://eksisozluk.com" + scrape.Attr(topicLink, "href")

		titleAndCount := scrape.Text(topicNode)

		countIndex := strings.LastIndex(titleAndCount, " ")

		topic.Title = strings.TrimSpace(titleAndCount[0:countIndex])
		countString := titleAndCount[countIndex:]

		topicCountInt, _ := strconv.Atoi(strings.TrimSpace(countString))
		topic.Count = int64(topicCountInt)

		topicList = append(topicList, topic)
	}

	return topicList
}

func (s Scraper) GetDEBE() []Debe {

	resp, err := http.Get("https://eksisozluk.com/debe")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	topicListNode, _ := scrape.Find(root, s.contentMatcher)

	debeTopics := s.getTopics(topicListNode)

	debeList := make([]Debe, 0)

	for _, t := range debeTopics {
		currentDebe := Debe{}
		currentDebe.DebeTopic = t
		entryList := s.getEntries(t.Link)
		currentDebe.DebeEntry = entryList[0]
		debeList = append(debeList, currentDebe)
	}

	return debeList
}
