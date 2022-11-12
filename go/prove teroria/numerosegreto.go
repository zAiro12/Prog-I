package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var x string
	fmt.Print("Inserisci il numero segreto tra '#': ")
	fmt.Scan(&x)
	y, err := secret(x)
	if err {
		fmt.Println("Il doppio del numero segreto inserito è:", y)
	} else {
		fmt.Println("Errore nell'input")
	}
}

func secret(s string) (int, bool) {
	var t string
	for _, r := range s {
		if unicode.IsDigit(r) {
			t += string(r)
		} else if r != '#' {
			return 0, false
		}
	}
	x, err := strconv.Atoi(t)
	if err != nil {
		return 0, false
	}
	return 2 * x, true
}
