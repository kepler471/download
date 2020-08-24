package main

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

//handleURL gets the response from a given URL and ensures it is HTML
func handleURL(URL string) *http.Response {
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		URL = "http://" + URL
	}
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalf("error fetching URL: %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("response status code was %d\n", resp.StatusCode)
	}
	ctype := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ctype, "text/html") {
		log.Fatalf("response content type was %s not text/html\n", ctype)
	}
	return resp
}

// getLinks finds all <a> tags, and returns their href attribute values
func getLinks(resp *http.Response) (links []string) {
	tokens := html.NewTokenizer(resp.Body)
	defer resp.Body.Close()
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

// getInfo creates and populates a File struct, given a link to a file
func getInfo(URL string, link string) File {
	u, _ := url.Parse(URL)

	var f = File{
		Type: FileType(*t),
		Name: path.Base(link),
		URL:  u.Scheme + "://" + path.Join(u.Host, path.Dir(u.Path), link),
	}
	resp, err := http.Head(f.URL)
	if err != nil {
		log.Fatalf("error fetching file header: %v @ %v\n", f.Name, f.URL)
	}
	defer resp.Body.Close()
	f.Size = uint64(resp.ContentLength)
	return f
}
