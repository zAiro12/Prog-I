package main

import "fmt"

func main() {
	var n, g int
	var utente string
	fmt.Print("Come ti chiami?: ")
	fmt.Scan(&utente)
	fmt.Print(utente, " inserisci un numero intero grande: ")
	fmt.Scan(&n)
	g = n % 10
	fmt.Print("L'unità del numero inserito è: ", g, ", gg ", utente, "!\n")
}
