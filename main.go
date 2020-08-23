package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// "http://www.gatsby.ucl.ac.uk/teaching/courses/ml1-2016.html"

/*
TODO create a `file` type, something along these lines

type File struct {
	Type	FileType
	Data    string
	URL     string
	Size	int

type FileType string

*/

func main() {
	filetype := os.Args[1]
	fmt.Println(filetype)
	for _, URL := range os.Args[2:] {
		if !strings.HasPrefix(URL, "http://") {
			URL = "http://" + URL
		}
		fmt.Println(URL)
		resp, err := http.Get(URL)
		// check URL fetched correctly
		if err != nil {
			log.Fatalf("error fetching URL: %v\n", err)
		}
		//check response status code
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("response status code was %d\n", resp.StatusCode)
		}
		//check response content type
		ctype := resp.Header.Get("Content-Type")
		if !strings.HasPrefix(ctype, "text/html") {
			log.Fatalf("response content type was %s not text/html\n", ctype)
		}
		// get links from all anchor tag references
		links := getLinks(resp.Body)
		_ = resp.Body.Close()
		// download links that match the given filetype
		var files []string
		for _, link := range links {
			if strings.HasSuffix(link, "."+filetype) {
				files = append(files, link)
				fmt.Printf("Link found: %v\n", link)
			}
		}
		if len(links) == 0 {
			fmt.Println("No links found")
		} else {
			fmt.Printf("%v links found", len(links))
		}
	}
}

func getLinks(body io.Reader) (links []string) {
	tokens := html.NewTokenizer(body)
	for {
		find := tokens.Next()
		switch find {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := tokens.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}

		}
	}
}

func downloadFile() {

}
