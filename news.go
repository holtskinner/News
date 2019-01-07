package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SitemapIndex top level xml
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

// News Individual News Stories
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>Keywords"`
	Locations []string `xml:"url>loc"`
}

func newsFetcher() {

	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		fmt.Printf("%s", Location)
	}

	grades := make(map[string]float32)
	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67

}
