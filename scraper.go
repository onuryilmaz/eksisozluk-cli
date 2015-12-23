package main

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var entryMatcher func(n *html.Node) bool
var authorMatcher func(n *html.Node) bool
var dateMatcher func(n *html.Node) bool
var entryListMatcher func(n *html.Node) bool
var topicListMatcher func(n *html.Node) bool
var contentMatcher func(n *html.Node) bool

func init() {

	entryListMatcher = func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "id"), "entry-list")
	}

	entryMatcher = func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "class"), "content")
	}

	authorMatcher = func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "class"), "entry-author")
	}

	dateMatcher = func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "class"), "entry-date")
	}

	topicListMatcher = func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "class"), "topic-list")
	}

	contentMatcher = func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "id"), "content-body")
	}

}

// GetEntries gets string for topic with parameters and returns a list of entries
func GetEntries(text string, parameter Parameter) []Entry {

	baseURL := "https://eksisozluk.com/?q=" + url.QueryEscape(text)

	resp, err := http.Get(baseURL)
	if err != nil {
		panic(err)
	}

	redirectedURL := resp.Request.URL.String()

	entryList := make([]Entry, 0)
	startPage := parameter.PageNumber

	for parameter.Limit > len(entryList) {

		paginationURL := redirectedURL + "?p=" + strconv.Itoa(startPage)

		if parameter.Sukela {
			paginationURL = paginationURL + "&a=nice"
		}

		additionalEntryList := getEntries(paginationURL)
		if len(additionalEntryList) == 0 {
			break
		}
		if len(entryList)+len(additionalEntryList) > parameter.Limit {
			entryList = append(entryList, additionalEntryList[0:(parameter.Limit-len(entryList))]...)
		} else {
			entryList = append(entryList, additionalEntryList...)
		}

		startPage = startPage + 1
	}

	return entryList
}


// GetPopularTopics gets parameters and returns a list of topics
func GetPopularTopics(parameter Parameter) []Topic {

	baseURL := "https://eksisozluk.com/basliklar/populer"

	topicList := make([]Topic, 0)
	startPage := parameter.PageNumber

	for parameter.Limit > len(topicList) {
		paginationURL := baseURL + "?p=" + strconv.Itoa(startPage)
		additionalTopicList := getTopics(paginationURL)
		if len(additionalTopicList) == 0 {
			break
		}
		if len(topicList)+len(additionalTopicList) > parameter.Limit {
			topicList = append(topicList, additionalTopicList[0:(parameter.Limit-len(topicList))]...)
		} else {
			topicList = append(topicList, additionalTopicList...)
		}

		startPage = startPage + 1
	}

	return topicList

}

// GetDEBE gets parameters and returns a list of Debe
func GetDEBE(parameter Parameter) []Debe {

	debeTopics := getTopics("https://eksisozluk.com/debe")

	debeList := make([]Debe, 0)

	for _, t := range debeTopics {
		if len(debeList) >= parameter.Limit {
			break
		}
		t.Count = 1 // Auto-correct count to 1 since only one entry is provided in DEBE
		currentDebe := Debe{}
		currentDebe.DebeTopic = t
		entryList := getEntries(t.Link)
		currentDebe.DebeEntry = entryList[0]
		debeList = append(debeList, currentDebe)
	}

	return debeList
}

func getEntries(eksiURL string) []Entry {

	log.Println("URL to check: " + eksiURL)
	resp, err := http.Get(eksiURL)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	entryList := make([]Entry, 0)

	entryListNode, found := scrape.Find(root, entryListMatcher)

	if found == false {
		return entryList
	}

	scrapedEntries := scrape.FindAll(entryListNode, entryMatcher)

	if len(scrapedEntries) == 0 {
		return entryList
	}
	for _, scrappedEntry := range scrapedEntries {
		authorNode, authorCheck := scrape.Find(scrappedEntry.Parent, authorMatcher)
		dateNode, dateCheck := scrape.Find(scrappedEntry.Parent, dateMatcher)

		entry := Entry{}
		entry.Text = scrape.Text(scrappedEntry)

		if authorCheck {
			entry.Author = scrape.Text(authorNode)
		}
		if dateCheck {
			idDate := scrape.Text(dateNode)
			splitted := strings.SplitAfterN(idDate, " ", 2)
			entry.ID = strings.TrimSpace(splitted[0])
			entry.Date = strings.TrimSpace(splitted[1])
		}
		entryList = append(entryList, entry)
	}

	return entryList
}

func getTopics(topicURL string) []Topic {

	log.Println("Checking URL: " + topicURL)
	topicList := make([]Topic, 0)

	resp, err := http.Get(topicURL)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	contentNode, _ := scrape.Find(root, contentMatcher)

	topicListNode, _ := scrape.Find(contentNode, topicListMatcher)

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
