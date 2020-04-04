package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {

	// Instantiate default collector
	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(1),
	)

	thUrl := "https://www.thehindu.com/news/international/brazils-bolsonaro-questions-coronavirus-deaths-says-sorry-some-will-die/article31189205.ece"

	c.OnHTML(`#content-body-14269002-31189205`, func(e *colly.HTMLElement) {
		pText := e.ChildText("p")

		fmt.Println("", pText)

		name := e.Name
		fmt.Println(name)

		text := e.Text
		fmt.Println(text)

		// Goquery selection of the HTMLElement is in e.DOM
		goquerySelection := e.DOM

		//goquerySelection.Find("a").Siblings().Each(func(i int, s *goquery.Selection) {
		//	fmt.Printf("%d, Sibling text: %s\n", i, s.Text())
		//})

		goquerySelectionTxt := goquerySelection.Text()
		fmt.Println(goquerySelectionTxt)

		nodes := goquerySelection.Nodes
		fmt.Println(nodes)

		DataAtom := nodes[0].DataAtom.String()
		fmt.Println(DataAtom)

		// Example Goquery usage
		span := goquerySelection.Find(" span").Children().Text()
		fmt.Println(span)

		ahref := e.ChildText("a[href]")
		fmt.Println(ahref)





		/*e.ForEach("p", func(_ int, el *colly.HTMLElement) {
			elCT := el.ChildText("")
			elHref := el.Attr("href")

			fmt.Println("", elCT)
			fmt.Println("", elHref)

		})*/
	})


	////////////////////////////////////////////////////////////////////////////////

	c.OnHTML(`div[class=article-topics-container]`, func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fmt.Println("", link)

		elCT := e.ChildText("")
		elHref := e.Attr("href")

		fmt.Println("", elCT)
		fmt.Println("", elHref)

	})

	////////////////////////////////////////////////////////////////////////////////



	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", "links")
	})

	// Start scraping on thUrl
	c.Visit(thUrl)
}
