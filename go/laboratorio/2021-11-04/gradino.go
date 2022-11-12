package main

import "fmt"

func main() {
	var x int
	// var b rune
	var i1, i2, imax, cont int
	for {
		fmt.Print("input: ")
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}
		if x != i1 {
			i2 = i1
			i1 = x
			imax = cont
			cont = 1
		} else {
			cont++
		}
	}
}
