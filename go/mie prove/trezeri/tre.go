package main

import "fmt"

func main() {
	var x int

	fmt.Print("inserisci un numero che termina con almeno 000: ")
	fmt.Scan(&x)

	if x%1000 != 0 || x == 0 {
		fmt.Println("no")
	} else {
		fmt.Println("si")
	}
}
