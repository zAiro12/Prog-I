package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	parole := os.Args[1:]
	for i := 0; i < len(parole); i++ {
		fmt.Print(TrasformaParola(parole[i], i), " ")
	}
	fmt.Println()
}

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
	if posizione%2 == 0 {
		for pos, val := range parola {
			if pos%2 == 0 {
				parolaTrasformata += strings.ToUpper(string(val))
			} else {
				parolaTrasformata += string(val)
			}
		}
	} else {
		for pos, val := range parola {
			if pos%2 != 0 {
				parolaTrasformata += strings.ToUpper(string(val))
			} else {
				parolaTrasformata += string(val)
			}
		}
	}
	return
}
