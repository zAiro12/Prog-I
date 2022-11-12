package main

import "fmt"

func main() {
	var s string
	fmt.Print("Inserisci un numero in binario: ")
	fmt.Scan(&s)
	x, ok := binToString(s)
	if ok {
		fmt.Println("Il valore è", x)
	} else {
		fmt.Println("Errore nell'input")
	}
}

func binToString(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	var v int
	for _, r := range s {
		switch r {
		case '0':
			v *= 2
		case '1':
			v = 2*v + 1
		default:
			return 0, false
		}
	}
	return v, true
}
