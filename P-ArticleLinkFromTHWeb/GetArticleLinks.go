package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)

func mainnnnnnn() {

	// Instantiate default collector
	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(1),
	)

	thUrl := "https://www.thehindu.com/news/international/"

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		scappingUrl := thUrl
		articleUrl := ".ece"

		if !strings.Contains(link, scappingUrl) {
			fmt.Println("Not valid url, So it is not considered ")
			return
		}

		if !strings.Contains(link, articleUrl) {
			fmt.Println("Not Article URL, So it is not considered ")
			return
		}

		// Print link
		fmt.Println(link)

		fileName := "International_Articles_Link.txt"

		f, err := os.OpenFile(fileName,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(link + "\n"); err != nil {
			log.Println(err)
		}

		// Visit link found on page
		//e.Request.Visit(link)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", "links")
	})

	// Start scraping on thUrl
	c.Visit(thUrl)
}
