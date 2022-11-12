package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	magica := genStringa()
	for {
		var attempt string
		fmt.Scan(&attempt)
		gpg, gps := check(magica, attempt)
		fmt.Printf("Magica: %s \t GPG: %d \t GPS: %d \n", magica, gpg, gps)
		if gpg == 5 {
			break
		}
	}
	fmt.Println("HAI INDOVINATO!!!")
}

func genStringa() string {
	s := ""
	for i := 0; i < 5; i++ {
		s += string('A' + rand.Intn(6))
	}
	return s
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
