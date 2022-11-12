package main

import "fmt"

func main() {
	var a, b, c, m float64

	fmt.Print("Di quanti numeri vuoi fare la media? 2 o 3? ")
	fmt.Scan(&m)

	if m == 2 {
		fmt.Print("Inserisci il primo numero: ")
		fmt.Scan(&a)
		fmt.Print("Inserisci il secondo numero: ")
		fmt.Scan(&b)

		m = (a + b) / 2
		fmt.Println("la media dei due numeri inseriti è", m)
	} else {
		fmt.Print("Inserisci il primo numero: ")
		fmt.Scan(&a)
		fmt.Print("Inserisci il secondo numero: ")
		fmt.Scan(&b)
		fmt.Print("Inserisci il terzo numero: ")
		fmt.Scan(&c)

		m = (a + b + c) / 3
		fmt.Println("la media dei due numeri inseriti è", m)
	}

}
