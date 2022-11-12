package main

import "fmt"

func main() {
	var a, b float64

	fmt.Scan(&a, &b)

	fmt.Print("somma = ", a+b, "\n", "differenza = ", a-b, "\n", "prodotto = ", a*b, "\n", "quoziente = ", a/b)

}
