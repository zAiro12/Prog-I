package main

import (
	"fmt"
)

type data struct {
	g, m, a int
}

func main() {
	var s data
	fmt.Print("inserisci una data: ")
	fmt.Scan(&s.g, &s.m, &s.a)
	Calc(s.g, s.m, s.a)

}

func Calc(g, m, a int) int {
	var ris int
	for i := 1970; i < a; i++ {
		if Bis(i) {
			ris += 366
		} else {
			ris += 365
		}
	}
	for i := 1; i < m; i++ {
		if i == 11 || i == 4 || i == 6 || i == 9 {
			ris += 30
		} else if i == 2 {
			if Bis(a) {
				ris += 29
			} else {
				ris += 28
			}
		} else {
			ris += 31
		}
	}
	ris += g
	return ris
}

func Bis(aa int) bool {
	if aa >= 4 && aa <= 1580 && aa%4 == 0 {
		return true
	} else if (aa > 1580) && (aa%100 == 0) {
		if aa%400 == 0 {
			return true
		} else {
			return false
		}
	} else if aa > 1580 && aa%4 == 0 {
		return true
	} else {
		return false
	}
}
