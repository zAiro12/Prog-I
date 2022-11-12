package main

import "fmt"

func main() {
	var s byte
	var mem string
	for {
		fmt.Print("inserisci una lettera, se vuoi finire metti '.': ")
		fmt.Scanf("%c", &s)
		mem += string(s)
		if s == '.' {
			break
		}
	}
	for _, a := range mem {
		if a >= 'a' && a <= 'z' {
			fmt.Printf("%c è la %d^a \n", a, a-'a'+1)
		} else if a == '.' {
			fmt.Println("bye")
		} else {
			fmt.Println("altro")
		}
	}

}
