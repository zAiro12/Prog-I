package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64
	fmt.Print("x e y del primo punto: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("x e y del secondo punto: ")
	fmt.Scan(&x2, &y2)

	distanza := math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
	fmt.Println("La distanza tra i due punti è:", distanza)
}
