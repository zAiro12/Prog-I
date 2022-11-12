package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, resto, mem int
	var deci int
	fmt.Print("inserisci un numero in binario da convertire in decimale: ")
	fmt.Scan(&x)
	mem = x
	for i := 0; resto < 2; i++ {
		resto = x % 10
		if resto == 0 || resto == 1 {
		} else {
			fmt.Print("Hai inserito ", resto, ". In binario i caratteri ammessi sono solo 0 e 1 \n")
			fmt.Println("Non hai inserito un numero in binario")
			break
		}
		y = x / 10
		resto = resto * int(math.Pow(2, float64(i)))
		deci += resto
		//fmt.Println("x:", x, "y:", y, "deci:", deci, "resto:", resto)
		x = y
		resto = 0
		if y == 0 {
			fmt.Println(mem, "in base bianria equivale a", deci, "in base decimale")
			break
		}
	}

}
