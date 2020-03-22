package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		//colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
		//colly.AllowedDomains("thehindu.com", "thehindu.com/thehindu.com/news/national/"),
		colly.AllowedDomains("thehindu.com", "www.thehindu.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting 1", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
		fmt.Println("Something went wrong:", err)
	})

	fmt.Println("Ashwani Start")
	// Start scraping on https://hackerspaces.org
	c.Visit("thehindu.com/sport/cricket/janani-vrinda-named-in-international-panel-of-icc-development-umpires/article31102416.ece")
	fmt.Println("Ashwani END")
}
