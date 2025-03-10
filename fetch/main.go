package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = fmt.Sprintf("http://%s", url)
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)

		}
		//b, err := io.ReadAll(resp.Body)
		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Problems reading %s", url)
			os.Exit(1)
		}
		/*
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s\n", b)
		*/
		fmt.Printf("Read %d lines\n", b)
		fmt.Println("Statuscode: ", resp.Status)
	}
}
