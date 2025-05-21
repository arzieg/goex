// rendezvous behaviour.
//  beim rendezvous behaviour werden die channels nicht gebuffert. D.h. lesen und schreiben sind voneinander abhängig

package main

import (
	"fmt"
	"time"
)

// T is a minimal structure
type T struct {
	i byte
	b bool
}

func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)}
	ch <- t

	// ein blocking sending. D.h. alle t.b sind false. Dies liegt daran, dass es unbuffered ist.
	// im Main wird go send ausgeführt, die function send sendet false zurück und der empfangene
	// channel wartet auch solang, bis es einen Wert zurückerhält (blocking sending). Dies ist
	// zu diesem Zeitpunkt false. Sender sendet, receiver receivs und erst dann ist das SENDEN beendet!

	// Anders im Fall mit buffered channel (a la email). Hier wird senden und empfangen entkoppelt, d.h.
	// t.b wird auf true gesetzt und der empfangende channal sieht diesen Zustand

	t.b = true // UNSAFE
}

func main() {
	vs := make([]T, 5)
	//ch := make(chan *T) // unbuffered channel
	ch := make(chan *T, 5) // buffered channel with 5 go-routines

	for i := range vs {
		go send(i, ch)
	}

	time.Sleep(1 * time.Second)

	//copy quickly
	for i := range vs {
		vs[i] = *<-ch
	}

	// print
	for _, v := range vs {
		fmt.Println(v)
	}

}
