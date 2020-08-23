package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type File struct {
	Type FileType
	Name string
	URL  string
	Size int64
}

type FileType string

var t = flag.String("t", "pdf", "specify file type")
var y = flag.Bool("y", false, "assume yes for download confirmation")

// TODO add flag to suppress output

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
		var totalDownloadSize int64
		var files []File
		for _, link := range links {
			if strings.HasSuffix(link, "."+*t) {
				u, _ := url.Parse(URL)
				// get URL for file and store
				var f = File{
					Type: FileType(*t),
					Name: path.Base(link),
					URL:  u.Scheme + "://" + path.Join(u.Host, path.Dir(u.Path), link),
				}
				// get file information
				resp, err := http.Head(f.URL)
				if err != nil {
					log.Fatalf("error fetching file: %v @ %v\n", f.Name, f.URL)
				}
				f.Size = resp.ContentLength
				totalDownloadSize += f.Size
				fmt.Printf("%v, size: %v: %v\n", f.Name, f.Size, f.URL)
				if *y {
					downloadFile(f)
				} else {
					files = append(files, f)
				}
				_ = resp.Body.Close()
			}
		}

		if len(files) == 0 {
			fmt.Print("No links found")
		} else {
			fmt.Printf("\n%v %v files found\n", len(files), *t)
			fmt.Printf("Total download size: %v\n", totalDownloadSize)
			fmt.Print("Would you like to download all [y/N]: ")
			var input string
			_, err := fmt.Scanln(&input)
			input = strings.Trim(input, " ")
			fmt.Println(input)
			if err != nil {
				log.Println("invalid user input")
			}
			switch input {
			case "y", "Y":
				for _, f := range files {
					downloadFile(f)
				}
			default:
				fmt.Println("Aborted.")
			}
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

func downloadFile(f File) {
	fmt.Printf("~ downloading %v\n", f.Name)
}
