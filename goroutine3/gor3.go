package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {

	var r result

	start := time.Now()
	ticker := time.NewTicker(1 * time.Second).C
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {
		r = result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Microsecond)
		r = result{url, nil, t}
		resp.Body.Close()
	}

	for {
		select {
		case ch <- r:
			return
		case <-ticker:
			log.Println("tick", r)
		}
	}
}

func first(ctx context.Context, urls []string) (*result, error) {
	// es kann sein, dass sender sendet, aber receiver noch nichts empfängt, weil noch nicht gestartet.
	// das kann zum Hochlaufen führen. mit len(urls) kann man die Kapazität begrenzen
	results := make(chan result, len(urls)) // buffer to avoid leaking
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	for _, url := range urls {
		go get(ctx, url, results)

	}
	select {
	case r := <-results: // when return also defer cancel() will run
		return &r, nil
	case <-ctx.Done(): // handle timeout from context above, if that happens. I have to handle this
		return nil, ctx.Err()
	}
}

func main() {

	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
		"http://localhost:8000/slow",
	}

	r, _ := first(context.Background(), list)

	if r.err != nil {
		log.Printf("%-20s %s\n", r.url, r.err)
	} else {
		log.Printf("%-20s %s\n", r.url, r.latency)

	}

	time.Sleep(9 * time.Second)
	log.Println("quit anyway...", runtime.NumGoroutine(), "still running")

}
