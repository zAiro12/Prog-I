package main

import (
	"fmt"
)

func main() {
	var k int
	var s, mem string
	fmt.Print("chiave: ")
	fmt.Scan(&k)
	fmt.Print("caratteri da cifrare: ")
	fmt.Scan(&s)
	for _, a := range s {
		if a >= 'a' && a <= 'z' {
			a = rune(((int(a)+k)-'a')%26 + 'a')
			mem += string(int(a))
		}

	}
	fmt.Println(mem, "è il testo cifrato")
}
