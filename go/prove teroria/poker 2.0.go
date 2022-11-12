package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"time"
)

type carta struct {
	valore, seme int
}

func main() {
	var n int
	var mazzo, mano []carta
	rand.Seed(time.Now().UnixNano())
	mazzo = genmazzo()
	//quente persone giocano
	fmt.Print("Quante persone giocano? ")
	fmt.Scan(&n)
	mani := make([][]carta, 0)
	nomi := make([]string, 0)
	//salvo i nomi dei giocatori
	for i := 0; i < n; i++ {
		var nome string
		fmt.Printf("Come si chiama il giocatore %d ", i+1)
		fmt.Scan(&nome)
		nomi = append(nomi, nome)
	}

	//generazione mani
	for i := 0; i < n; i++ {
		mano, mazzo = genmano(mazzo) //shuffle mazzo e gen mano
		mani = append(mani, mano)
		mani[i] = mano
	}
	//cambio carte
	for i := 0; i < n; i++ {
		var c int //num carte da cambiare
		pulisi()
		fmt.Printf("Turno di %s \n", nomi[i])
		time.Sleep(2 * time.Second)
		vedicarte(mani[i])
		fmt.Printf("%s, quante carte vuoi cambiare? ", nomi[i])
		fmt.Scan(&c)
		if c != 0 {
			fmt.Print("Quali carte vuoi cambiare? ")
			for k := 0; k < c; k++ {
				var h int
				fmt.Scan(&h)
				mani[i], mazzo = cambioCarta(mani[i], mazzo, h-1)
			}
		}
		vedicarte(mani[i])
		time.Sleep(4 * time.Second)
		pulisi()
		fmt.Println("Puoi passare il dispositivo")
		for j := 4; j > 0; j-- {
			fmt.Print(j, "...\n")
			time.Sleep(1 * time.Second)
		}
	}
	mappe := []map[int]int{}
	mappes := []map[int]int{}
	for i := 0; i < n; i++ {
		mappe = append(mappe, creaMappaVal(mani[i]))
		mappes = append(mappes, creaMappaSeme(mani[i]))
	}

	fmt.Println(colore(mappes))

}

//funzioni
func pulisi() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func creacarta(c carta) string {
	valoreNome := [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	semeNome := [...]string{"\u2660", "\u2663", "\u2666", "\u2665"}
	return valoreNome[c.valore] + semeNome[c.seme]
}

func genmazzo() []carta {
	var mazzo []carta
	for s := 0; s < 4; s++ {
		for v := 0; v < 13; v++ {
			c := carta{valore: v, seme: s}
			mazzo = append(mazzo, c)
		}
	}
	return mazzo
}

func genmano(mazzo []carta) ([]carta, []carta) {
	rand.Shuffle(len(mazzo), func(i, j int) { mazzo[i], mazzo[j] = mazzo[j], mazzo[i] })
	mano := mazzo[:5]
	mazzo = mazzo[5:]
	return mano, mazzo
}

func cambioCarta(mano []carta, mazzo []carta, i int) ([]carta, []carta) {
	switch i {
	case 0:
		mano[0] = mazzo[0]
		mazzo = mazzo[1:]
		break
	case 1:
		mano[1] = mazzo[0]
		mazzo = mazzo[1:]
		break
	case 2:
		mano[2] = mazzo[0]
		mazzo = mazzo[1:]
		break
	case 3:
		mano[3] = mazzo[0]
		mazzo = mazzo[1:]
		break
	case 4:
		mano[4] = mazzo[0]
		mazzo = mazzo[1:]
		break
	}
	return mano, mazzo
}

func vedicarte(mano []carta) {
	fmt.Print("La tua mano è: ")
	for _, a := range mano {
		fmt.Print(creacarta(a), "\t")

	}
	fmt.Println()
}

func creaMappaVal(mani []carta) map[int]int {
	mappa := map[int]int{}
	for i := 0; i < 5; i++ {
		mappa[mani[i].valore]++
	}
	return mappa
}
func creaMappaSeme(mani []carta) map[int]int {
	mappa := map[int]int{}
	for i := 0; i < 5; i++ {
		mappa[mani[i].seme]++
	}
	return mappa
}

func check() {

}

func alta(mappe []map[int]int) []int {
	var ris = []int{}
	for i := 0; i < len(mappe); i++ {
		max := 0
		for key := range mappe[i] {
			if key > max {
				max = key
			}
		}
		ris = append(ris, max)
	}
	return ris
}

func coppia(mappe []map[int]int) []int {
	var ris = []int{}
	for i := 0; i < len(mappe); i++ {
		max := -1
		for key, valore := range mappe[i] {
			if valore == 2 && key > max {
				max = key
			}
		}
		ris = append(ris, max)
	}
	return ris
}

func doppiacoppia(mappe []map[int]int) []int {
	var ris = []int{}
	for i := 0; i < len(mappe); i++ {
		max := -1
		flag := false
		for key, valore := range mappe[i] {
			if valore == 2 && key > max {
				max = key
			}
		}
		if max != -1 {
			for key, valore := range mappe[i] {
				if valore == 2 && key != max {
					flag = true
					if key > max {
						max = key
					}
				}
			}
			if !flag {
				max = -1
			}
		}
		ris = append(ris, max)
	}
	return ris
}

func tris(mappe []map[int]int) []int {
	var ris = []int{}
	for i := 0; i < len(mappe); i++ {
		max := -1
		for key, valore := range mappe[i] {
			if valore == 3 && key > max {
				max = key
			}
		}
		ris = append(ris, max)
	}
	return ris
}

func colore(sem []map[int]int) []int {
	var ris = []int{}
	for i := 0; i < len(sem); i++ {
		max := -1
		for key, valore := range sem[i] {
			if valore == 5 && key > max {
				max = key
			}
		}
		ris = append(ris, max)
	}
	return ris
}

func full(mappe []map[int]int) []int {
	var ris = []int{}
	for i := 0; i < len(mappe); i++ {
		max := -1
		for key, valore := range mappe[i] {
			if valore == 3 && key > max {
				max = key
				if valore == 1 && key != max {
					ris = append(ris, max)
				}
			}
		}
		ris = append(ris, max)
	}
	return ris
}

func scala(mappe []map[int]int) []int {
	var ris = []int{}
	for i := 0; i < len(mappe); i++ {
		max := -1
		c := []int{}
		for key := range mappe[i] {
			c = append(c, key)
		}
		sort.Ints(c)
		fmt.Println(c)
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

func poker(mappe []map[int]int) []int {
	var ris = []int{}
	for i := 0; i < len(mappe); i++ {
		max := -1
		for key, valore := range mappe[i] {
			if valore == 4 && key > max {
				max = key
			}
		}
		ris = append(ris, max)
	}
	return ris
}

func scalacol(mappe, mappes []map[int]int) {
	var ris = []int{}
	for i := 0; i < len(mappe); i++ {
		max := -1
		c := []int{}
		for key := range mappe[i] {
			c = append(c, key)
		}
		sort.Ints(c)
		fmt.Println(c)
		flag := false
		for key, valore := range mappes[i] {

		}
		if len(c) == 5 {
			for i := 0; i < 4; i++ {
				if c[i] == c[i+1]-1 {
					max = c[i+1]
				}
			}
		}
		ris = append(ris, max)
	}
}

func scalarel(mappe []map[int]int) {

}
