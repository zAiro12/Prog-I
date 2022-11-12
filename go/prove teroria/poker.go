package main

import (
	"fmt"
)

func main() {
	//var x int
	var carte [52]int
	var mano [4]string
	var x int
	// fmt.Print("Inserisci un numero da 1 a 52: ")
	// fmt.Scan(&x)
	//rand.Seed(time.Now().Unix())
	for h := 0; h <= 4; h++ {
		fmt.Scan(&x)
		mano[h] = genCarte(x)
	}
	/*for j := 0; j < 5; j++ {
		for {
			x = rand.Intn(53)
			if x >= 1 && x <= 52 {
				break
			}
		}

		for i := 1; i <= 52; i++ {
			if carte[i] != x {
				c = genCarte(x)
			}
		}
		mano[j] = c
	}*/
	for h := 1; h <= 5; h++ {
		fmt.Print(carte[h])
	}
}

func genCarte(x int) string {
	var s string
	switch x % 13 {
	case 0:
		s = "re"
	case 1:
		s = "asso"
	case 2:
		s = "2"
	case 3:
		s = "3"
	case 4:
		s = "4"
	case 5:
		s = "5"
	case 6:
		s = "6"
	case 7:
		s = "7"
	case 8:
		s = "8"
	case 9:
		s = "9"
	case 10:
		s = "10"
	case 11:
		s = "jack"
	case 12:
		s = "donna"

	}
	switch x / 14 {
	case 0:
		s += " picche"
	case 1:
		s += " fiori"
	case 2:
		s += " cuori"
	case 3:
		s += " quadri"
	}
	return s
}
