package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, resto, mem int
	var bin string
	fmt.Print("inserisci un numero intero da convertire in biario: ")
	fmt.Scan(&x)
	mem = x
	y = 1
	if x >= 0 {
		for y != 0 {
			resto = x % 2
			y = x / 2
			bin = fmt.Sprint(resto) + bin
			x = y
			resto = 0
		}
		fmt.Println(mem, "in base decimale equivale a", bin, "in base binaria")
	} else {
		var i, n int
		x = int(math.Abs(float64(x)))

		for i = 0; n < x; i++ {
			n = int(math.Pow(2, float64(i)))
		}
		n = int(math.Pow(2, float64(i)))
		x = n - x
		x = x % int(math.Pow(10, float64(i)))
		fmt.Println(x, n, i)

		for y != 0 {
			resto = x % 2
			y = x / 2
			bin = fmt.Sprint(resto) + bin
			x = y
			resto = 0
		}
		fmt.Println(mem, "in base decimale equivale a", bin, "in complemento a due")
	}

}
