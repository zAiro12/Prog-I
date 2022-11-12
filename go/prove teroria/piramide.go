package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci lunghezza della base del triangolo: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := n; j > 0; j-- {
			if j > i {
				fmt.Print(" ")
			} else {
				fmt.Print("* ")
			}

		}
		fmt.Println()
	}
}
