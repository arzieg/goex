package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type xkcd struct {
	Num        int    `json:"num"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no file given")
		os.Exit(1)
	}

	f := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "no search terms")
		os.Exit(1)
	}

	var (
		items []xkcd
		terms []string
		input io.ReadCloser
		cnt   int
		err   error
	)

	// get search terms
	if input, err = os.Open(f); err != nil {
		fmt.Fprintf(os.Stderr, "bad file name: %s\n", err)
		os.Exit(1)
	}

	// decode file
	if err = json.NewDecoder(input).Decode(&items); err != nil {
		fmt.Fprintf(os.Stderr, "bad json: %s \n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", len(items))
	// get search term
	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

	// search
outer:
	for _, item := range items {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Transcript)

		for _, term := range terms {
			if !strings.Contains(title, term) && !strings.Contains(transcript, term) {
				continue outer
			}
		}

		fmt.Printf("https://xkcd.com/%d/ %s/%s/%s %q\n",
			item.Num, item.Month, item.Day, item.Year, item.Title,
		)
		cnt++
	}

	fmt.Fprintf(os.Stderr, "found %d comics\n", cnt)
}
