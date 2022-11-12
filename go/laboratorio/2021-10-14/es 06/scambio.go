package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("val1 e val2 (int):")
	_, err := fmt.Scan(&a, &b)

	fmt.Println("errore:", err)
	fmt.Println("prima dello scambio:", a, b)
	c := a
	a = b
	b = c
	fmt.Println("dopo dello scambio:", a, b)
}
