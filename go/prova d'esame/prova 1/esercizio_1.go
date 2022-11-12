package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var s1, s2, s3 rune
	var a, b, c int
	fmt.Sscanf(os.Args[1], "%c%d%c%dx%c%dx^2=0", &s1, &c, &s2, &b, &s3, &a)
	if s1 == '-' {
		c = -c
	}
	if s2 == '-' {
		b = -b
	}
	if s3 == '-' {
		a = -a
	}
	delta := float64((b * b) - (4 * a * c))
	ris1 := (float64(-b) + math.Sqrt(float64(delta))) / float64(2*a)
	ris2 := (float64(-b) - math.Sqrt(float64(delta))) / float64(2*a)

	if ris1 == ris2 {
		fmt.Printf("Esiste un’unica soluzione reale: %.2f\n", ris1)
	}
	if (ris1 != ris2) && (math.IsNaN(ris1) == false) {
		fmt.Printf("Ci sono due soluzioni reali: %.2f e %.2f \n", ris1, ris2)
	}
	if math.IsNaN(ris1) && math.IsNaN(ris2) {
		fmt.Println("Non ci sono soluzioni reali")
	}
}
