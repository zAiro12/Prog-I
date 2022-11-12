package main

import "fmt"

func main() {
	var voto int

	fmt.Print("Voto: ")
	fmt.Scan(&voto)

	if voto >= 0 && voto <= 100 {
		if voto <= 100 && voto >= 90 {
			fmt.Println("A")
		} else if voto <= 89 && voto >= 80 {
			fmt.Println("B")
		} else if voto <= 79 && voto >= 70 {
			fmt.Println("C")
		} else if voto <= 69 && voto >= 60 {
			fmt.Println("D")
		} else if voto <= 59 && voto >= 0 {
			fmt.Println("F")
		}
	} else {
		fmt.Println("voto non valido")
	}
}
