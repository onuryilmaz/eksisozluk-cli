package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestScraper(t *testing.T) {

	Convey("GetPopularTopics return hot topics", t, func() {
		p := Parameter{}
		p.Limit = 3
		topics := GetPopularTopics( p)
		So(len(topics), ShouldEqual, 3)
	})

	Convey("GetEntries return entries for topic", t, func() {
		p := Parameter{}
		p.Limit = 10
		entries := GetEntries("golang", p)
		So(len(entries), ShouldEqual, 10)
	})

}
