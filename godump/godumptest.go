package main

import (
	"fmt"
	"os"

	"github.com/goforj/godump"
)

type Profile struct {
	Age   int
	Email string
}

type User struct {
	Name    string
	Profile Profile
}

func main() {
	user := User{
		Name: "Alice",
		Profile: Profile{
			Age:   30,
			Email: "alice@example.com",
		},
	}

	// Pretty-print to stdout
	godump.Dump(user)

	// Dump and exit
	//godump.Dd(user) // this will print the dump and exit the program

	// Get dump as string
	output := godump.DumpStr(user)
	fmt.Println("str", output)

	// HTML for web UI output
	html := godump.DumpHTML(user)
	fmt.Println("html", html)

	// Write to any io.Writer (e.g. file, buffer, logger)
	godump.Fdump(os.Stderr, user)
}
