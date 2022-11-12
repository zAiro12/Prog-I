package main

import "fmt"

func main() {
	var reddito int
	var coniugato bool
	const ali10, ali25, sca32, sca64 = 10 / 100, 25 / 100, 32000, 64000
	fmt.Print("inserisci il rettido e se sei coniugato (t/f): ")
	fmt.Scan(&reddito, &coniugato)
	if reddito >= 0 && reddito <= sca32 && !coniugato {
		fmt.Println(float64(reddito) * ali10)
	} else if reddito > sca32 && !coniugato {
		fmt.Println(float64(reddito) * ali25)
	} else if reddito >= 0 && reddito <= sca64 && !coniugato {
		fmt.Println(float64(reddito) * ali10)
	} else {
		fmt.Println(float64(reddito) * ali25)
	}
}
