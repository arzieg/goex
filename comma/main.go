package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func bufToComma(s string) string {
	// pre-allocate sufficient buffer capacity
	var buf bytes.Buffer
	buf.Grow(len(s) + len(s)/3)

	count := 0

	// Loop from the last character (least-significant digit) to the first.
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
		count++
		// Insert a comma after every 3 digits if there are more digits remaining.
		if count == 3 && i != 0 {
			buf.WriteByte(',')
			count = 0
		}
	}

	// The bytes in the buffer are in reverse order.
	// Reverse the underlying byte slice to obtain the correct order.
	b := buf.Bytes()
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: comma <argument1>")
		os.Exit(1)
	}

	fmt.Println("Zahl mit comma über comma: ", comma(os.Args[1]))
	fmt.Println("Zahl mit comma über buffer:", bufToComma(os.Args[1]))

}
