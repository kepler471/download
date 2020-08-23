package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// "http://www.gatsby.ucl.ac.uk/teaching/courses/ml1-2016.html"

func main() {
	filetype := os.Args[1]
	fmt.Println(filetype)
	for _, URL := range os.Args[2:] {
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

		page, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		re, err := regexp.Compile("body")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Regex error: #{err}\n")
		}
	}
}

func getLinks(body io.Reader) []string {
	var links []string
	tok := html.NewTokenizer(resp.Body)
}

func downloadFile() {

}
