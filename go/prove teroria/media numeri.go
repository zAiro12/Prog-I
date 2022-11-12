package main

import "fmt"

func main() {
	var x, n int
	sum := 0
	for {
		fmt.Print("Inserisci un numero (0 per terminare): ")
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		sum += x
		n++
	}

	media := float64(sum) / float64(n)
	fmt.Println("La media di", n, "numeri è", media)
}
