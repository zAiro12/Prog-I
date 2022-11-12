package main

import "fmt"

func main() {
	var f func(float64) float64
	f = func(x float64) float64 { return 3*x*x + 5*x + 3 }
	a := integra(f, 7, 12, 100)
	fmt.Println(a)
}


