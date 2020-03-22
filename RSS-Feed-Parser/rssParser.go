package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.thehindu.com/news/states/feeder/default.rss")
	fmt.Println(feed.Items)
	fmt.Println("Ashwani")
}
