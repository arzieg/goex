package main

import (
	"fmt"
)

func main() {
	//	defer printStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panic x=0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/*
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
*/
