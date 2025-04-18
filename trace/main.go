package main

import (
	"log"
	"time"
)

// Example defer on-entry and on-exit

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // dont forget the extra parentheses
	time.Sleep(10 * time.Second)      // simulate slow operation
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) } // return anonymous function
}

func main() {
	bigSlowOperation()
}
