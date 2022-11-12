package main

import (
	"fmt"
)

func main() {
	var in int
	var n []int
	fmt.Scan(&in)
	for in != -1 {
		n = append(n, in)
		fmt.Scan(&in)

	}
	fmt.Println(sommaRic(n))
}

func sommaRic(n []int) int {
	if len(n) == 0 {
		return 0
	}
	return n[0] + sommaRic(n[1:])

}
