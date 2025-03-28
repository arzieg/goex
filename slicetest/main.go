package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
		fmt.Printf("i=%d  j=%d\n", i, j)
	}
}

func main() {
	months := [...]string{
		1:  "Jan",
		2:  "Feb",
		3:  "Mar",
		4:  "Apr",
		5:  "Mai",
		6:  "Jun",
		7:  "Jul",
		8:  "Aug",
		9:  "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dec",
	}

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Printf("Q2: Typ: %T, Werte:%v, Capacity: %d, Len: %d\n", Q2, Q2, cap(Q2), len(Q2))
	fmt.Printf("summer: Typ: %T, Werte:%v, Capacity: %d, Len: %d\n", summer, summer, cap(summer), len(summer))

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println("Reverse: ", a)

	s := []int{0, 1, 2, 3, 4, 5}
	// rotate s left by two positions
	reverse(s[:2])
	fmt.Println("s=", s)
	reverse(s[2:])
	fmt.Println("s=", s)
	reverse(s)
	fmt.Println("Reverse by 2: ", s)
}
