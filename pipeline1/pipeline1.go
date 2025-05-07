package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x < 101; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//squarer
	go func() {
		for x := range naturals {
			squares <- x * x
			time.Sleep(100 * time.Millisecond)
		}
		close(squares)
	}()

	// print in main
	for x := range squares {
		//fmt.Println(<-naturals)
		fmt.Println(x)
	}
}
