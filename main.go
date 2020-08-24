package main

import (
	"flag"
	"fmt"
	"github.com/dustin/go-humanize"
	"strings"
)

var (
	t = flag.String("t", "pdf", "specify file type")
	l = flag.Bool("l", false, "list files only, overrides all bool flags")
	y = flag.Bool("y", false, "assume yes for download confirmation")
	U = flag.Bool("u", false, "display URL with file information")
	//o = flag.String("o", "", "specify filepath for download")
)

func main() {
	flag.Parse()
	// TODO want to handle URL endings, eg .html, .htm ...
	// TODO handle even when 0 urls provided with better feedback
	// TODO provide help, default for no args
	// TODO replace logs with errors
	var (
		files     []File
		totalSize uint64
	)
	for _, URL := range flag.Args() {
		if !*l {
			fmt.Println(URL)
		}
		links := getLinks(handleURL(URL))
		for _, link := range links {
			if strings.HasSuffix(link, "."+*t) {
				f := getInfo(URL, link)
				if !*l {
					fmt.Printf("size: %v, %v,\t%v\n", humanize.Bytes(f.Size), f.Name, f.URL)
				} else {
					fmt.Println(f.URL)
				}
				files = append(files, f)
				totalSize += f.Size
			}
		}
	}
	// UI, suppressed by -l flag
	if !*l && len(files) > 0 {
		fmt.Printf("\n%v %v files found\n", len(files), *t)
		fmt.Printf("Total download size: %v\n", humanize.Bytes(totalSize))
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
					f.Download()
				}
			default:
				fmt.Println("Aborted.")
			}
		} else {
			for _, f := range files {
				f.Download()
			}
		}
	}
}
