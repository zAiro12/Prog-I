package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type data struct {
	g, m, a int
}

func main() {
	var data []data
	for i := 1; i < len(os.Args); i++ {

		d, ok := strData(os.Args[i])
		if ok {
			data = append(data, d)

		} else {
			fmt.Println("Errore: data sbagliata")
		}
	}
	d, i := prima(data)
	fmt.Println("la data più piccola è:", d, "in posizione:", i)
}

func strData(s string) (data, bool) {
	p1 := strings.IndexRune(s, '/')
	if p1 < 0 {
		return data{}, false
	}
	p2 := strings.IndexRune(s, '/')
	if p2 < 0 {
		return data{}, false
	}
	p2 += p1 + 1
	g, err := strconv.Atoi(s[:p1])
	if err != nil {
		return data{}, false
	}
	m, err := strconv.Atoi(s[p1+1 : p2])
	if err != nil {
		return data{}, false
	}
	a, err := strconv.Atoi(s[p2+1:])
	if err != nil {
		return data{}, false
	}
	return data{g, m, a}, true
}

func vienePrima(d1, d2 data) bool {
	if d1.a != d2.a {
		return d1.a < d2.a
	}
	if d1.m != d2.m {
		return d1.m < d2.m
	}
	return d1.g < d2.g
}

func prima(d []data) (data, int) {
	cand := d[0]
	indCand := 0
	for i := 1; i < len(d); i++ {
		if vienePrima(d[i], cand) || d[i] == cand {
			cand = d[i]
			indCand = i
		}

	}
	return cand, indCand
}
