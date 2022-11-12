package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	l, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])
	c := os.Args[3][0]

	if l < 0 || n < 0 {
		return
	} else {
		for i := 0; i < n; i++ {
			fmt.Println(DrawSegment(c, i*(l-1), l))
			for j := 0; j < l-1; j++ {
				fmt.Println(DrawPoint(c, (i+1)*(l-1)))
			}
		}
	}
}

func DrawPoint(c byte, k int) string {
	var s string
	for i := 0; i < k; i++ {
		s += " "
	}
	s += string(c)
	return s
}

func DrawSegment(c byte, k, l int) string {
	var s string
	for i := 0; i < k; i++ {
		s += " "
	}
	for i := 0; i < l; i++ {
		s += string(c)
	}
	return s
}
