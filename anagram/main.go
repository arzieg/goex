package main

import (
	"fmt"
	"os"
)

func anagram(s1, s2 string) bool {
	runes := []rune(s2)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return s1 == string(runes)

}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: anagram <string1> <string2>")
		os.Exit(1)
	}

	fmt.Printf("%s und %s sind ein Anagram? %t\n", os.Args[1], os.Args[2], anagram(os.Args[1], os.Args[2]))
}
