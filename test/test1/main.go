package main

import (
	"fmt"
	"math/rand"
	"test1/word"
	"time"
)

func main() {
	fmt.Printf("is test palindrome = %t\n", word.IsPalindrome("test"))

	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	p := word.RandomPalindrome(rng)
	fmt.Printf("RandomPalindrome = %s\n", p)

}
