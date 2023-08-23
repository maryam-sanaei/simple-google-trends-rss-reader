package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Create structs based on rss structure in
// https://trends.google.com/trends/trendingsearches/daily/rss?geo=AU (or another location)
// and based on what we need to get. (some items not needed)

type RSS struct {

	// we have one channel to every rss, so we just point to it
	Channel *Channel `xml:"channel"`
}

type Channel struct {

	// every channels can have many items
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	PubDate   string `xml:"pubDate"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadLineLink string `xml:"news_item_url"`
}

var url string

func main() {

	var r RSS

	// you can specify rss url with -url switch, default would be us
	flag.StringVar(&url, "url", "https://trends.google.com/trends/trendingsearches/daily/rss?geo=US", "RSS url")
	flag.Parse()

	data := readGoogleTrends()
	err := xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Here all google search trends for taday")
	fmt.Println("------------------------------------------------------------")
	for i := range r.Channel.ItemList {
		rank := i + 1
		fmt.Println("#", rank)
		fmt.Println("Title:", r.Channel.ItemList[i].Title)
		fmt.Println("Link:", r.Channel.ItemList[i].Link)
		fmt.Println("Publish Date:", r.Channel.ItemList[i].PubDate)
		fmt.Println()
		fmt.Println("Head Lines:")
		for j := range r.Channel.ItemList[i].NewsItems {
			fmt.Println("Title:", r.Channel.ItemList[i].NewsItems[j].Headline)
			fmt.Println("Ralated Link:", r.Channel.ItemList[i].NewsItems[j].HeadLineLink)
			fmt.Println()
		}
		fmt.Println("--------------------------------------------------------")

	}

}
func getGoogleTrends() *http.Response {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}
