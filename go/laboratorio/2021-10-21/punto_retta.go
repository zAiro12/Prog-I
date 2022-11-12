package main

import "fmt"

func main() {
	var x, y, m, q float64
	const kekw = 10e-6
	fmt.Print("inserisci x e y di un punto: ")
	fmt.Scan(&x, &y)
	fmt.Print("inserisci m e q di una retta: ")
	fmt.Scan(&m, &q)
	r := (((m * x) + q) - y) / kekw

	if r > 0.0 {
		fmt.Println("sopra")
	} else if r < 0.0 {
		fmt.Println("sotto")
	} else {
		fmt.Println("sulla retta")
	}
}
