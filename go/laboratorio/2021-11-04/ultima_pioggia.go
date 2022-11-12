package main

import (
	"fmt"
	"io"
)

func main() {
	var x, g int
	var controllo bool = false
	for i := 0; ; i++ {
		_, err := fmt.Scan(&x)
		if x == 0 {
			g = i
			controllo = true
		}
		if err == io.EOF {
			if x == 0 {
				g--
			}
			break
		}
	}
	if controllo {
		fmt.Println("output:", g)
	} else {
		fmt.Println("non ha mai piovuto")
	}

}
