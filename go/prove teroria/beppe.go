package main

import (
	"fmt"
	"math"
	"math/rand"
)

type tentativo struct {
	s        string
	gpg, gps int
}

func main() {
	var gpg, gps int
	var tentativi []tentativo
	tutte := sequenze(5)
	for {
		alcune := filtra(tutte, tentativi)
		fmt.Println(tutte[rand.Intn(len(tutte))])
		t := alcune[rand.Intn(len(alcune))]
		fmt.Scan(&gpg, &gps)
		if gpg == 5 {
			break
		}
		tentativi = append(tentativi, tentativo{t, gpg, gps})
	}
	fmt.Println("HO VINTO, YEE!!")
}

func sequenze(n int) []string {
	if n == 0 {
		return []string{""}
	}
	seq := sequenze(n - 1)
	var risultato []string
	for _, s := range seq {
		for car := range []rune{'A', 'B', 'C', 'D', 'E', 'F'} {
			t := s + string(car)
			risultato = append(risultato, t)
		}
	}
	return risultato
}

func filtra(seq []string, tent []tentativo) []string {
	var res []string
	for _, s := range seq {
		ok := true
		for _, t := range tent {
			gpg, gps := check(s, t.s)
			if gpg != t.gpg || gps != t.gps {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, s)
		}

	}
	return res
}

func check(magic, attempt string) (gpg int, gps int) {
	var magicMap, attemptMap map[byte]int
	magicMap = make(map[byte]int)
	attemptMap = make(map[byte]int)
	for i := 0; i < 5; i++ {
		if magic[i] == attempt[i] {
			gpg++
		} else {
			magicMap[magic[i]]++
			attemptMap[magic[i]]++
		}
	}
	for r, occ := range magicMap {
		gps += int(math.Min(float64(occ), float64(attempt[r])))
	}
	return
}
