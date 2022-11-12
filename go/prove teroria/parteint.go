package main

import "fmt"

func main() {
	var i int
	var f float64
	fmt.Scan(&f, &i)
	if verifica(f, i) {
		fmt.Println("si")
	} else {
		fmt.Println("no")
	}
}

func verifica(f float64, i int) bool {
	return int(f) == i
}
