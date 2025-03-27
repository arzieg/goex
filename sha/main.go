package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var (
	// commandline flags
	t int
)

func init() {
	flag.IntVar(&t, "t", 2, "Type 2,3 or 5")
}

func main() {

	var hash []byte
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: sha -t [2,3,5] <string>")
		os.Exit(1)
	}

	switch t {
	case 3:
		c := sha512.Sum384([]byte(args[0]))
		hash = c[:]
	case 5:
		c := sha512.Sum512([]byte(args[0]))
		hash = c[:]
	default:
		c := sha256.Sum256([]byte(args[0]))
		hash = c[:]
	}
	fmt.Printf("Hash Type: %T\nHash value:  %x\n", hash, hash)
}
