package main

import (
	"Publisher/P-CategoryFromTHWeb/model"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/url"
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

	//links := &[]model.Category{}
	links := make(map[string]model.Category)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// Print link
		// fmt.Println(link)

		fileName := "THLinks5.txt"

		f, err := os.OpenFile(fileName,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(link + "\n"); err != nil {
			log.Println(err)
		}
		pathSpliter(link, links)

		// Visit link found on page
		e.Request.Visit(link)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", links)
	})

	fmt.Println(links)

	// Start scraping on https://www.thehindu.com/news
	c.Visit("https://www.thehindu.com/news")
}

//func pathSpliter(rawURL string, links *[]model.Category) {
func pathSpliter(rawURL string, links map[string]model.Category) {
	//rawURL := "https://www.thehindu.cm/life-and-style/homes-and-gardens/"
	scappingUrl := "https://www.thehindu.com/"
	articleUrl := ".ece"

	fmt.Println(rawURL)

	if !strings.Contains(rawURL, scappingUrl) {
		fmt.Println("Not valid url, So it is not considered ")
		return
	}

	if strings.Contains(rawURL, articleUrl) {
		fmt.Println("Article URL, So it is not considered ")
		return
	}

	url, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	path := url.Path
	uriSegments := strings.Split(path, "/")
	pathLen := len(uriSegments)
	pathCount := 0

	catName := ""
	subCatName := ""

	for i := 0; i < pathLen; i++ {
		segment := uriSegments[i]

		//subCategory := model.SubCategory{}
		//subSubCategory := model.SubSubCategory{}

		if segment != "" {

			pathCount++
			if segment == "news" {
				fmt.Println("")
			}
			if pathCount == 1 {
				fmt.Println("Category ", segment)
				catName = segment
				if category, ok := links[catName]; ok {
					category.CategoryName = catName
					links[catName] = category
				} else {
					category := model.Category{}
					category.CategoryName = catName
					links[catName] = category
				}

			}
			if pathCount == 2 {
				fmt.Println("Sub-Category ", segment)

				if(catName == "profile") {
					log.Println("")
				}

				if category, ok := links[catName]; ok {

					subCategory := model.SubCategory{}

					if category.SubCategories == nil {
						category.SubCategories = make(map[string]model.SubCategory)
					} else {
						subCategory = category.SubCategories[segment]
					}
					subCategory.SubCategoryName = segment
					category.SubCategories[segment] = subCategory
					links[catName] = category

					fmt.Println("Sub-Category ", segment)
					subCatName = segment
				}
			}
			if pathCount >= 3 {
				if(catName == "profile") {
					log.Println("")
				}
				fmt.Println("Sub-Sub-Category ", segment)
				if category, ok := links[catName]; ok {
						if subCategory, ok := category.SubCategories[subCatName]; ok {
							subSubCategory := model.SubSubCategory{}

							if subCategory.SubSubCategories == nil {
								subCategory.SubSubCategories = make(map[string]model.SubSubCategory)
							} else {
								subSubCategory = subCategory.SubSubCategories[segment]
							}
							subSubCategory.SubSubCategoryName = segment
							subCategory.SubSubCategories[segment] = subSubCategory

							category.SubCategories[subCatName] = subCategory
						}
						links[catName] = category

				}
			}
		}
	}

}
