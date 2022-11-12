package main

import "fmt"

func main() {
	var x int
	fmt.Print("Inserici un numero dispari: ")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Println("Il numero non è dispari")
	} else {
		//prima riga
		for i := 0; i < x; i++ {
			fmt.Print("*")
		}
		fmt.Println()
		//per il resto
		for i, t := 1, x-1; t >= 2; {
			for j := 0; j < t/2; j++ {
				fmt.Print("*")
			}
			for j := 0; j < i; j++ {
				fmt.Print(" ")
			}
			for j := 0; j < t/2; j++ {
				fmt.Print("*")
			}
			fmt.Println()
			t -= 2
			i += 2
		}

	}
}
