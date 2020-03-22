package main

import (
	"fmt"
	"strings"
)


func main() {
	//rawURL := "http://example.com/index.php/news/local/metro/crime_is_up"
	rawURL := "https://www.thehindu.cm/life-and-style/homes-and-gardens/"
	scappingUrl := "https://www.thehindu.com/"
	if(!strings.Contains(rawURL, scappingUrl)) {
		fmt.Println("NO segment ")
		return
	}
	rawURL = strings.ReplaceAll(rawURL, scappingUrl , "")

	uriSegments := strings.Split(rawURL, "/")

	fmt.Println(uriSegments) // count starts from 1

	for _, val := range uriSegments {
		fmt.Println("segment ", val)
	}

	for i:=0; i< len(uriSegments); i++ {
		fmt.Println("segment New ", uriSegments[i])
	}
}
