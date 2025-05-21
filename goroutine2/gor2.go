package main

import (
	"context"
	"log"
	"net/http"
	"time"
	//"golang.org/x/net/context"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if resp, err := http.DefaultClient.Do(req); err != nil {
		ch <- result{url, err, 0}

	} else {
		t := time.Since(start).Round(time.Microsecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}
func main() {
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
		"http://localhost:8000/slow",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	for _, url := range list {
		go get(ctx, url, results)
	}

	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)

		}
	}

}
