package main

import "fmt"

func main() {
	var a float64

	fmt.Print("Inserisci un numero con la virgola: ")
	fmt.Scan(&a)
	b := int(a)
	a = a - float64(b)
	fmt.Println("La parte decimale del numero inserito è", a)
}
