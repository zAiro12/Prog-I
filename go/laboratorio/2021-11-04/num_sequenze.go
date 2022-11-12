package main

import "fmt"

func main() {
	var s rune
	var c, cmax int
	for {
		fmt.Print("inserisci 1 o 0 (2 per fermare): ")
		fmt.Scan(&s)
		if s == 0 {
			c++
			if c >= cmax {
				cmax = c
			}
		} else if s == 1 {
			c = 0
		} else if s == 2 {
			break
		}
	}
	fmt.Println("output:", cmax)

}
