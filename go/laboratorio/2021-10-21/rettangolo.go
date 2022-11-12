package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 int
	fmt.Print("P1: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("P2: ")
	fmt.Scan(&x2, &y2)
	fmt.Print("P3: ")
	fmt.Scan(&x3, &y3)

	if x3 < int(math.Min(float64(x1), float64(x2))) || x3 > int(math.Max(float64(x1), float64(x2))) || y3 < int(math.Min(float64(y1), float64(y2))) || y3 > int(math.Max(float64(y1), float64(y2))) {
		fmt.Println("esterno")
	} else {
		fmt.Println("interno")
	}
}
