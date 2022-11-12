package main

import (
	"fmt"
)

func main() {
	var i, n int

	fmt.Print("Quante volte vuoi che ti saluto? ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Println("A scemoooo!!!")
	}
}
