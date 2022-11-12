package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var a, b, c int
	var s1, s2, s3 rune
	var ris1, ris2 float64
	input := os.Args[1]
	if input[0] == '-' {
		fmt.Sscanf(input, "%c%dx^2%c%dx%c%d=0", &s1, &a, &s2, &b, &s3, &c)
	} else {
		fmt.Sscanf(input, "%dx^2%c%dx%c%d=0", &a, &s2, &b, &s3, &c)
	}
	if s2 == '-' {
		b = -b
	}
	if s3 == '-' {
		c = -c
	}
	dis := b*b - (4 * a * c)

	ris1 = (float64(-b) + math.Sqrt(float64(dis))) / float64(2*a)
	ris2 = (float64(-b) - math.Sqrt(float64(dis))) / float64(2*a)

	if dis < 0 {
		fmt.Println("Non ci sono soluzioni reali")
	} else if dis == 0 {
		fmt.Printf("Esiste un'unica soluzione reale: %.3f\n", ris1)
		controlla(ris1, ris2)
	} else {
		fmt.Printf("Esistono due soluzioni reali: %.3f e %.3f\n", ris1, ris2)
		controlla(ris1, ris2)
	}
}

func controlla(ris1 float64, ris2 float64) {
	soglia := os.Args[2]
	epsilon := os.Args[3]
	s, _ := strconv.Atoi(soglia)
	e, _ := strconv.Atoi(epsilon)
	if ris1 > float64(s) && ris1-float64(s) > float64(e) {
		fmt.Printf("La soluzione %.3f è maggiore della soglia\n", ris1)
	} else if ris2 > float64(s) && ris2-float64(s) > float64(e) {
		fmt.Printf("La soluzione %.3f è maggiore della soglia\n", ris2)
	}
}
