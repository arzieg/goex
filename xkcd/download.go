package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type xkcdjson struct {
	Num        int    `json:"num"`
	Month      int    `json:"month"`
	Day        string `json:"day"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func getOne(i int) []byte {

	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't read: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "skipping %d: got %d\n", i, resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid body: %s\n", err)
		os.Exit(1)
	}
	return body
}

func main() {

	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		fails  int
		data   []byte
	)

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer output.Close()
	}

	fmt.Println("[")
	defer fmt.Println("]")

	for i := 1; fails < 2; i++ {
		if data = getOne(i); data == nil {
			fails++
			continue
		}

		if cnt > 0 {
			fmt.Fprintf(output, ",")
		}

		_, err = io.Copy(output, bytes.NewBuffer(data))

		if err != nil {
			fmt.Fprintf(os.Stderr, "stopped: %s\n", err)
			os.Exit(1)
		}

		fails = 0
		cnt++
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
}
