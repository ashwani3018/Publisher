package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(1),
	)

	// On every a element which has href attribute call callback

	links := []string{}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//c.OnHTML("div", func(e *colly.HTMLElement) {
		//link := e.Attr("div")
		link := e.Attr("href")

		/*e.ForEach("body", func(_ int, el *colly.HTMLElement) {
			fmt.Println("Name :: "+el.Name)
			fmt.Println("Text :: "+el.Text)
		})*/
		// Print link
		fmt.Println(link)
		links = append(links, link)
		//linkByte := [] byte(link)

		fileName := "THLinks.txt"

		f, err := os.OpenFile(fileName,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(link+"\n"); err != nil {
			log.Println(err)
		}


		// Visit link found on page
		e.Request.Visit(link)
	})


	// Start scraping on https://en.wikipedia.org
	//c.Visit("https://en.wikipedia.org/")
	c.Visit("https://www.thehindu.com/news")
}
