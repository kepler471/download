package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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

var (
	t = flag.String("t", "pdf", "specify file type")
	y = flag.Bool("y", false, "assume yes for download confirmation")
	// TODO flag to output only matching files
	totalDownloadSize int64
	files             []File
)

func main() {
	flag.Parse()
	// TODO want to handle URL endings, eg .html, .htm ...
	// TODO handle even when 0 urls provided with better feedback
	// TODO provide help
	// TODO without any arguments, show the help
	for _, URL := range flag.Args() {
		resp := handleURL(URL)
		links := getLinks(resp.Body)
		_ = resp.Body.Close()
		for _, link := range links {
			if strings.HasSuffix(link, "."+*t) {
				files = append(files, getInfo(URL, link))
			}
		}
	}
	// UI
	fmt.Printf("\n%v %v files found\n", len(files), *t)
	if len(files) > 0 {
		// TODO better size formatting here and in getInfo
		fmt.Printf("Total download size: %v k\n", totalDownloadSize/1000)
		if !*y {
			fmt.Print("Would you like to download all [y/N]: ")
			var input string
			_, err := fmt.Scanln(&input)
			input = strings.Trim(input, " ")
			if err != nil {
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
	f.Size = resp.ContentLength
	totalDownloadSize += f.Size
	fmt.Printf("%v,\tsize: %v,\t%v\n", f.Name, f.Size, f.URL)
	_ = resp.Body.Close()
	return f
}

// downloadFile uses io.Copy to save the response data from a file URL
func downloadFile(f File) {
	// TODO allow specifying save path
	file, err := os.Create(f.Name)
	if err != nil {
		log.Fatalf("could not create %v, %v\n", f.Name, err)
	}
	defer file.Close()
	data, err := http.Get(f.URL)
	if err != nil {
		log.Fatalf("error fetching file: %v\n", err)
	}
	_, err = io.Copy(file, data.Body)
	if err != nil {
		log.Fatalf("error writing to %v: %v\n", f.Name, err)
	}
}
