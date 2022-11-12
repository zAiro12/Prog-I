package main

import (
	"fmt"
	"sort"
)

type carta struct {
	valore, seme int
}

func main() {
	var mano1, mano2 []carta
	var mani [][]carta
	mano1 = append(mano1, carta{valore: 5, seme: 0})
	mani = append(mani, mano1)
	mano1 = append(mano1, carta{valore: 5, seme: 1})
	mani = append(mani, mano1)
	mano1 = append(mano1, carta{valore: 5, seme: 2})
	mani = append(mani, mano1)
	mano1 = append(mano1, carta{valore: 9, seme: 2})
	mani = append(mani, mano1)
	mano1 = append(mano1, carta{valore: 5, seme: 3})
	mani = append(mani, mano1)
	mani[0] = mano1

	mano2 = append(mano2, carta{valore: 9, seme: 0})
	mani = append(mani, mano2)
	mano2 = append(mano2, carta{valore: 8, seme: 0})
	mani = append(mani, mano2)
	mano2 = append(mano2, carta{valore: 11, seme: 0})
	mani = append(mani, mano2)
	mano2 = append(mano2, carta{valore: 10, seme: 0})
	mani = append(mani, mano2)
	mano2 = append(mano2, carta{valore: 12, seme: 0})
	mani = append(mani, mano2)
	mani[1] = mano2

	vedicarte(mano1)
	vedicarte(mano2)

	mappe := []map[int]int{}
	mappes := []map[int]int{}
	for i := 0; i < 2; i++ {
		mappe = append(mappe, creaMappa(mani[i]))
		mappes = append(mappes, creaMappaS(mani[i]))
	}

	fmt.Println(scalacol(mappe, mappes))
	/*if alta(mani[0]) > alta(mani[1]) {
		fmt.Println("vince 1")
	} else {
		fmt.Println("vince 2")
	}
	fmt.Print("coppia 1: ")
	fmt.Println(coppia(mani[0]))
	fmt.Print("coppia 2: ")
	fmt.Println(coppia(mani[1]))*/
	//fmt.Print("doppia coppia 1: ")
	//fmt.Println(doppiacoppia(mappe))
	//fmt.Print("doppia coppia 2: ")
	//fmt.Println(doppiacoppia(mappe))
	/*fmt.Print("tris 1: ")
	fmt.Println(tris(mani[0]))
	fmt.Print("tris 2: ")
	fmt.Println(tris(mani[1]))
	fmt.Print("full 1: ")
	fmt.Println(full(mani[0]))
	fmt.Print("full 2: ")
	fmt.Println(full(mani[1]))
	fmt.Print("scala 1: ")
	fmt.Println(scala(mani[0]))
	fmt.Print("scala 2: ")
	fmt.Println(scala(mani[1]))
	fmt.Print("poker 1: ")
	fmt.Println(poker(mani[0]))
	fmt.Print("poker 2: ")
	fmt.Println(poker(mani[1]))
	fmt.Print("scalacol 1: ")
	fmt.Println(scalacol(mani[0]))
	fmt.Print("scalacol 2: ")
	fmt.Println(scalacol(mani[1]))
	fmt.Print("scalarel 1: ")
	fmt.Println(scalarel(mani[0]))
	fmt.Print("scalarel 2: ")
	fmt.Println(scalarel(mani[1]))
	//fmt.Println(check(mani))*/
}

func vedicarte(mano []carta) {
	fmt.Print("La tua mano è: ")
	for _, a := range mano {
		fmt.Print(creacarta(a), "\t")

	}
	fmt.Println()
}

func creacarta(c carta) string {
	valoreNome := [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	semeNome := [...]string{"\u2660", "\u2663", "\u2666", "\u2665"}
	return valoreNome[c.valore] + semeNome[c.seme]

}

/* func check(mani [][]carta) string {
	var ciao []int
	for i := 0; i < 2; i++ {
		ok1, verifica1 := scalarel(mani[i])
		if ok1 {
			ciao = append(ciao, 10)
		}
		ok2, carta2, seme := scalacol(mani[i])
	}
	return "errore"
} */

func creaMappa(mani []carta) map[int]int {
	mappa := map[int]int{}
	for i := 0; i < 5; i++ {
		mappa[mani[i].valore]++
	}
	return mappa
}
func creaMappaS(mani []carta) map[int]int {
	mappa := map[int]int{}
	for i := 0; i < 5; i++ {
		mappa[mani[i].seme]++
	}
	return mappa
}

func scalacol(mappe, mappes []map[int]int) []int {
	var ris = []int{}
	for i := 0; i < len(mappe); i++ {
		max := -1
		c := []int{}
		for key := range mappe[i] {
			c = append(c, key)
		}
		sort.Ints(c)
		fmt.Println(c)
		// flag := false
		// for key, valore := range mappes[i] {

		// }
		if len(c) == 5 {
			for i := 0; i < 4; i++ {
				if c[i] == c[i+1]-1 {
					max = c[i+1]
				}
			}
		}
		ris = append(ris, max)
	}
	return ris
}

/* func alta(mani []carta) int {
	for j := 0; j < 5; j++ { //valore carta che sto controllando
		for k := 0; k < 5; k++ { //valore di carte da controllare
			for l := 0; l < 5; l++ {
				for m := 0; m < 5; m++ {
					for n := 0; n < 5; n++ {
						if mani[j].valore > mani[k].valore && mani[j].valore > mani[l].valore && mani[j].valore > mani[m].valore && mani[j].valore > mani[n].valore {
							return 1
						} else if mani[k].valore > mani[j].valore && mani[k].valore > mani[l].valore && mani[k].valore > mani[m].valore && mani[k].valore > mani[n].valore {
							return 2
						} else if mani[l].valore > mani[j].valore && mani[l].valore > mani[k].valore && mani[l].valore > mani[m].valore && mani[l].valore > mani[n].valore {
							return 3
						} else if mani[m].valore > mani[j].valore && mani[m].valore > mani[k].valore && mani[m].valore > mani[l].valore && mani[m].valore > mani[n].valore {
							return 4
						} else if mani[n].valore > mani[j].valore && mani[n].valore > mani[k].valore && mani[n].valore > mani[l].valore && mani[n].valore > mani[m].valore {
							return 5
						}

					}
				}
			}
		}
	}
	return 0
}

func coppia(mani []carta) (bool, int) {
	for j := 0; j < 5; j++ { //valore carta che sto controllando
		for k := 0; k < 5; k++ { //valore di carte da controllare
			if mani[j].valore == mani[k].valore {
				if mani[j].seme != mani[k].seme {
					return true, mani[j].valore
				}
			}
		}
	}

	return false, 0
}

func doppiacoppia(mani []carta) (bool, int) {
	for j := 0; j < 5; j++ { //valore carta che sto controllando
		for k := 0; k < 5; k++ { //valore della prima carta che confronto
			for l := 0; l < 5; l++ { //valore della seconda carta che prendo in considerazione
				for m := 0; m < 5; m++ { //valore della carta che confronto con la seconda carta
					if mani[j].valore == mani[k].valore && mani[l].valore == mani[m].valore {
						if mani[j].seme == mani[k].seme && mani[l].seme == mani[m].seme {
							return true, mani[j].valore
						}
					}
				}
			}
		}
	}
	return false, 0
}

func tris(mani []carta) (bool, int) {
	for j := 0; j < 5; j++ { //valore carta che sto controllando
		for k := 0; k < 5; k++ { //valore della prima carta che controllo
			for l := 0; l < 5; l++ { //valore della seconda carta che controllo
				if mani[j].valore == mani[k].valore && mani[k].valore == mani[l].valore && mani[j].valore == mani[l].valore {
					if mani[j].seme != mani[k].seme && mani[k].seme != mani[l].seme && mani[j].seme != mani[l].seme {
						return true, mani[j].valore
					}
				}
			}
		}
	}

	return false, 0
}

func scala(mani []carta) (bool, int) {
	ordina(mani)
	for i := 0; i < 4; i++ {
		if mani[i].valore == mani[i+1].valore-1 {
			return true, mani[4].valore
		}
	}
	return false, 0
}

func ordina(mani []carta) []carta {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if mani[j].valore < mani[i].valore {
				mani[j], mani[i] = mani[i], mani[j]
			}
		}
	}
	return mani
}

func full(mani []carta) (bool, int) {
	a, b := coppia(mani)
	c, d := tris(mani)
	if b != d {
		if a && c {
			return true, d
		}
	}
	return false, 0
}

func colore(mani []carta) (bool, int) {
	for i := 0; i < 4; i++ {
		if mani[i].seme == mani[i+1].seme {
			return true, mani[i].seme
		}
	}
	return false, 0
}

func poker(mani []carta) (bool, int) {
	for j := 0; j < 5; j++ { //valore carta che sto controllando
		for k := 0; k < 5; k++ { //valore di carte da controllare
			for l := 0; l < 5; l++ {
				for m := 0; m < 5; m++ {
					if mani[j].valore == mani[k].valore && mani[j].valore == mani[l].valore && mani[j].valore == mani[m].valore && mani[k].valore == mani[j].valore && mani[k].valore == mani[l].valore && mani[k].valore == mani[m].valore && mani[l].valore == mani[j].valore && mani[l].valore == mani[k].valore && mani[l].valore == mani[m].valore && mani[m].valore == mani[j].valore && mani[m].valore == mani[k].valore && mani[m].valore == mani[l].valore {
						if mani[j].seme != mani[k].seme && mani[j].seme != mani[l].seme && mani[j].seme != mani[m].seme && mani[k].seme != mani[j].seme && mani[k].seme != mani[l].seme && mani[k].seme != mani[m].seme && mani[l].seme != mani[j].seme && mani[l].seme != mani[k].seme && mani[l].seme != mani[m].seme && mani[m].seme != mani[j].seme && mani[m].seme != mani[k].seme && mani[m].seme != mani[l].seme {
							return true, mani[j].valore
						}
					}
				}
			}
		}
	}
	return false, 0
}

func scalacol(mano []carta) (bool, int, int) {
	a, b := scala(mano)
	if a {
		for j := 0; j < 5; j++ { //valore carta che sto controllando
			for k := 0; k < 5; k++ { //valore di carte da controllare
				for l := 0; l < 5; l++ {
					for m := 0; m < 5; m++ {
						for n := 0; n < 5; n++ {
							if mano[j].seme == mano[k].seme && mano[j].seme == mano[l].seme && mano[j].seme == mano[m].seme && mano[j] == mano[n] && mano[k].seme == mano[j].seme && mano[k].seme == mano[l].seme && mano[k].seme == mano[m].seme && mano[k] == mano[n] && mano[l].seme == mano[j].seme && mano[l].seme == mano[k].seme && mano[l].seme == mano[m].seme && mano[l] == mano[n] && mano[m].seme == mano[j].seme && mano[m].seme == mano[k].seme && mano[m].seme == mano[l].seme && mano[m] == mano[n] && mano[n].seme == mano[j].seme && mano[n].seme == mano[k].seme && mano[n].seme == mano[l].seme && mano[n] == mano[m] {
								return true, b, mano[j].seme
							}
						}
					}
				}
			}
		}
	}
	return false, 0, 0
}

func scalarel(mani []carta) (bool, int) {
	a, b, c := scalacol(mani)
	fmt.Println(a, b, c)
	if a {
		if b == 12 {
			return true, c
		}

	}
	return false, 0
} */
