package main

import "fmt"

func main() {
	var x, y float64
	const k = 180
	fmt.Print("angolo uno e due: ")
	fmt.Scan(&x, &y)

	fmt.Println("ampiezza terzo angolo:", k-x-y)
}
