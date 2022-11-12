package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type Punto struct {
	etichetta string
	x, y      float64
}

func main() {
	tutte := NuovoTragitto()
	lunghezza := Lunghezza(tutte)
	fmt.Printf("Lunghezza percorso: %.3f\n", lunghezza)
	
}

func NuovoTragitto() (tragitto []Punto) {
	for {
		var s string
		_, err := fmt.Scan(&s)
		if err == io.EOF {
			break
		}
		var nuovoPunto Punto
		split := strings.Split(s, ";")
		nuovoPunto.etichetta = split[0]
		nuovoPunto.x, _ = strconv.ParseFloat(split[1], 32)
		nuovoPunto.y, _ = strconv.ParseFloat(split[2], 32)
		tragitto = append(tragitto, nuovoPunto)
	}
	return tragitto
}

func Distanza(p1, p2 Punto) float64 {
	x := math.Pow(p1.x-p2.x, 2)
	y := math.Pow(p1.y-p2.y, 2)
	c := math.Sqrt(x + y)
	return c
}

func String(p Punto) string {
	var s string
	x := fmt.Sprintf("%v", p.x)
	y := fmt.Sprintf("%v", p.y)
	s = p.etichetta + " = (" + x + ", " + y + ")"
	return s
}

func Lunghezza(tragitto []Punto) float64 {
	c := 0.0
	for i := 0; i < len(tragitto)-1; i++ {
		c += Distanza(tragitto[i], tragitto[i+1])
	}
	return c
}
