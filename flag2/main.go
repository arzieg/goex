package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// commandline flags
	verbose   bool
	recursive bool
	method    string
	privkey   string
	pubkey    string
	cryptfile string
)

func init() {
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.BoolVar(&recursive, "r", false, "recursive scan of the directory")
	flag.StringVar(&method, "m", "c", "(c)rypt or (d)ecrypt")
	flag.StringVar(&privkey, "k", "", "private key")
	flag.StringVar(&pubkey, "p", "", "public key")
	flag.StringVar(&cryptfile, "f", "", "files to en-/decrypt")

}

func charInList(targetChar string, charList []string) bool {
	// Check if the target character is in the list
	found := false
	for _, char := range charList {
		if char == targetChar {
			found = true
			break
		}
	}
	return found
}

func main() {
	flag.Parse()

	// Validate that the input of method
	methodList := []string{"c", "C", "d", "D"}

	if len(method) != 1 {
		fmt.Println("Error: The -m flag must be a single character.")
		return
	}

	if !charInList(method, methodList) {
		fmt.Fprintf(os.Stderr, "Found %s as parameter for method, only one of %v is allowed\n", method, methodList)
		os.Exit(0)
	}

	fmt.Println("verbose:", verbose)
	fmt.Println("recursive:", recursive)
	fmt.Println("method:", method)
	fmt.Println("privkey:", privkey)
	fmt.Println("pubkey:", pubkey)
	fmt.Println("cryptfile:", cryptfile)
}
