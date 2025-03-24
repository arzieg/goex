package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < 256; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < 256; i++ {
		PopCountLoop(uint64(i))
	}
}
