package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type File struct {
	Type FileType
	Name string
	URL  string
	Size uint64
}

// Download uses io.Copy to save the response data from a file URL
func (f File) Download() {
	// see https://gobyexample.com/worker-pools for goroutines and channels
	file, err := os.Create(f.Name + ".tmp")
	if err != nil {
		log.Fatalf("could not create %v, %v\n", f.Name, err)
	}
	defer file.Close()
	data, err := http.Get(f.URL)
	if err != nil {
		log.Fatalf("error fetching file: %v\n", err)
	}
	defer data.Body.Close()
	counter := &WriteCounter{}
	_, err = io.Copy(file, io.TeeReader(data.Body, counter))
	if err != nil {
		log.Fatalf("error writing to %v: %v\n", f.Name, err)
	}
	fmt.Println()
	err = os.Rename(f.Name+".tmp", f.Name)
	if err != nil {
		log.Fatalf("error renaming temporary file: %v\n", f.Name)
	}
}

type FileType string

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 50))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}
