package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"golang.org/x/net/html"
)

func main() {
	filetype := os.Args[1]
	fmt.Println(filetype)
	for _, url := range os.Args[2:] {
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "url: #{err}\n")
		}
		re, err := regexp.Compile("body")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Regex error: #{err}\n")
		}
		page, err := ioutil.ReadAll(resp.Body)
		//resp.Body.Close()
			}
		}
	}
}

func getFileLinks(body io.Reader) []string {
	var links []string
	tok := html.NewTokenizer(resp.Body)
}

func downloadFile {

}