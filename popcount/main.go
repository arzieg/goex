package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var count byte
	for i := 0; i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}

func main() {
	var zahl uint64
	zahl = 127
	fmt.Printf("PopCount: %v\n", PopCount(zahl))
	fmt.Printf("PopCountLoop: %v\n", PopCountLoop(zahl))
}
