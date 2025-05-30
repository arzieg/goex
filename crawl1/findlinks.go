// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 240.

// Crawl1 crawls web links starting with the command-line arguments.
//
// This version quickly exhausts available file descriptors
// due to excessive concurrent calls to links.Extract.
//
// Also, it never terminates because the worklist is never closed.
package main

import (
	"fmt"
	"log"
	"os"

	//"gopl.io/ch5/links"
	links "links"
)

// enforce a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

// !+crawl
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // aquire a token
	list, err := links.Extract(url)

	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

// !+main
func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist

		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
