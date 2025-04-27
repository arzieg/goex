package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	count := 0
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*w += WordCounter(count)

	return count, nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*l += LineCounter(count)

	return count, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Printf("Charcount = %d\n", c)

	var w WordCounter
	w.Write([]byte("Das ist ein Text"))
	fmt.Printf("Wordcount: %d\n", w)

	var l LineCounter
	l.Write([]byte("Das ist ein Text\nmit\nmehreren Zeilenumbrüchen\n"))
	fmt.Printf("Linecount: %d\n", l)

	os.Exit(0)
}
