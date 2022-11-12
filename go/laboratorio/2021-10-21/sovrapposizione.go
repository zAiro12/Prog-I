package main

import "fmt"

func main() {
	var g1, g2, ini1, ini2, fin1, fin2 int
	fmt.Print("appuntamento 1 (gg, start, end): ")
	fmt.Scan(&g1, &ini1, &fin1)
	fmt.Print("appuntamento 2 (gg, start, end): ")
	fmt.Scan(&g2, &ini2, &fin2)

	if g1 < 1 || g2 < 1 || g1 > 31 || g2 > 31 || ini1 < 0 || ini2 < 0 || ini1 > 24 || ini2 > 24 || fin1 < 0 || fin2 < 0 || fin1 > 24 || fin2 > 24 {
		fmt.Println("errore")
	} else {
		if g1 == g2 || ini1 == ini2 || fin1 == fin2 {
			fmt.Println("si sovrappongono")
		} else {
			fmt.Println("non si sovrappongono")
		}
	}
}
