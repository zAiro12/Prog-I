package main

import "fmt"

func main() {
	const num = 2
	var x int
	fmt.Scan(&x)
	if x > 11 || x < 0 {
		fmt.Println("Valore non valido")
	} else {
		if x == num {
			fmt.Println("Hai indovinato!")
		} else {
			fmt.Println("Non hai indovinato")
		}
	}
}
