package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)

func main() {

	// Instantiate default collector
	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(1),
	)

	thUrl := "https://www.thehindu.com/news/international/"

	articleUrl := ".ece"

	c.OnHTML("section[class=feature-news]", func(e *colly.HTMLElement) {
		link := e.Attr("a")
		fmt.Println("", link)

		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			elCT := el.ChildText("")
			elHref := el.Attr("href")

			fmt.Println("", elCT)
			fmt.Println("", elHref)

			if !strings.Contains(elHref, articleUrl) {
				fmt.Println("Not Article URL, So it is not considered ")
				return
			}

			// Print link
			fmt.Println(elHref)

			fileName := "International_Articles_Link_On_Section.txt"

			f, err := os.OpenFile(fileName,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(elHref + "\n"); err != nil {
				log.Println(err)
			}

		})


	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", "links")
	})

	// Start scraping on thUrl
	c.Visit(thUrl)
}
