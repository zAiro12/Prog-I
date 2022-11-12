package main

import "fmt"

func main() {
	var voto30, voto110 float64

	fmt.Print("Dimmi il voto da convertire in 110esimi: ")
	fmt.Scan(&voto30)

	voto110 = (voto30 * 110) / 30

	fmt.Println("Il voto convertito equivale a:", voto110)
}
