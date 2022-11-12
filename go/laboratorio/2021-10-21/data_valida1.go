package main

import "fmt"

func main() {
	var gg, mm int
	fmt.Print("inserisci giorno e mese: ")
	fmt.Scan(&gg, &mm)
	if gg <= 0 || gg >= 31 || mm <= 0 || mm >= 12 {
		fmt.Println("data non valida")
	} else {
		fmt.Println("data valida")
	}
}
