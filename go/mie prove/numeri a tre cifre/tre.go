package main

import "fmt"

func main() {
	var x int

	fmt.Print("Inserisci un numero a tre cifre: ")
	fmt.Scan(&x)

	if x < 100 || x > 999 {
		print("Il numero che hai inserito è troppo grande o troppo piccolo")

	} else {
		a := (x % 10) + (x/10)%10 + (x/100)%10
		if a <= 10 {
			fmt.Println("la somma delle cifre del numero inserito è minore di 10")
		} else {
			fmt.Println("la somma delle cifre del numero inserito è maggiore di 10")
		}
	}

}
