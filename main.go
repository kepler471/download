package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"strings"
)

// "http://www.gatsby.ucl.ac.uk/teaching/courses/ml1-2016.html"

type File struct {
	Type FileType
	URL  string
	Size int
}

type FileType string

var t = flag.String("t", "pdf", "specify file type")
var y = flag.Bool("y", false, "assume yes for download confirmation")

func main() {
	flag.Parse()
	for _, URL := range flag.Args() {
		if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
			URL = "http://" + URL
		}
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
		// download links that match the given file type
		for _, link := range links {
			if strings.HasSuffix(link, "."+*t) {
				split := strings.Split(link, "/")
				link = split[cap(split)-1]
				f := File{
					Type: FileType(*t),
					URL:  strings.TrimSuffix(URL, ".html") + "/" + link,
					Size: 0,
				}
				if *y {
					downloadFile(f.URL)
				} else {
					fmt.Println("...user input here...")
				}
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

func downloadFile(link string) {
	fmt.Printf("~ Download placeholder for: %v\n", link)
}
