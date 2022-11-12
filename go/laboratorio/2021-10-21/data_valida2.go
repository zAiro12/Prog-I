package main

import "fmt"

func main() {
	var gg, mm int
	fmt.Print("inserisci giorno e mese: ")
	fmt.Scan(&gg, &mm)
	if mm == 1 || mm == 3 || mm == 5 || mm == 7 || mm == 8 || mm == 10 || mm == 12 {
		vergg(gg, 31)
	} else if mm == 4 || mm == 6 || mm == 9 || mm == 11 {
		vergg(gg, 30)
	} else if mm == 2 {
		vergg(gg, 28)
	} else {
		fmt.Println("data non valida")
	}
}

func vergg(gg int, k int) {
	if gg <= 0 || gg >= k {
		fmt.Println("data non valida")
	} else {
		fmt.Println("data valida")
	}
}
