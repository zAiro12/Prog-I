package main

import (
	"fmt"
)

func main() {
	var s string
	var c, cmax int
	for {
		fmt.Print("inserisci delle stringhe di numeri (0 per uscire): ")
		fmt.Scan(&s)
		if s == "0" {
			break
		}
		for _, a := range s {
			if a%2 == 0 {
				c++
				if c >= cmax {
					cmax = c
				}
			}
		}
		c = 0
	}
	fmt.Println(cmax)
}
