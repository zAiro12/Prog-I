package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserire il numero di righe: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		for j := 0; j < n*2-3; j++ { // Sbagliata la condizione
			fmt.Print(" ")
		}
		if n < 0 {
			fmt.Println()
			break
		}
		fmt.Print("*\n")
		n -= 2
	}
}
