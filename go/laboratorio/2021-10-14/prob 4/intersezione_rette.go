package main

import "fmt"

func main() {
	var m1, q1, m2, q2 float64

	fmt.Print("retta 1: m e q? ")
	fmt.Scan(&m1, &q1)
	fmt.Print("retta 2: m e q? ")
	fmt.Scan(&m2, &q2)

	fmt.Print("intersezione in (", (q1-q2)/(m2-m1), ",", m1*(q1-q2)/(m2-m1)+q1, ")")
}
