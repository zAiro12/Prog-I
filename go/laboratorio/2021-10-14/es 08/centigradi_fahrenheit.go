package main

import "fmt"

func main() {
	var g int
	var f float64
	const k = 32
	const j = 9. / 5

	fmt.Print("temperatura in centigradi? ")
	fmt.Scan(&g)

	f = (float64(g) * j) + k

	fmt.Println(g, "C =", f, "F")
}
