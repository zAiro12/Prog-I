package main

import "fmt"

func main() {
	var in int
	var n []int
	fmt.Scan(&in)
	for in != -1 {
		n = append(n, in)
		fmt.Scan(&in)

	}
	fmt.Println(ric(n))
}

func ric(n []int) bool {
	if len(n) == 0 {
		return false
	}
	if n[0] == 10 {
		return true
	}
	return ric(n[1:])
}
