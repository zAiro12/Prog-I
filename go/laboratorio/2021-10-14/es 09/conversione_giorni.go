package main

import "fmt"

func main() {
	const anno, sett = 365, 7
	var x int

	fmt.Println("numero di giorni da convertire?")
	fmt.Scan(&x)
	fmt.Println(x, "giorni =", x/anno, "anni +", x%anno/sett, "settimane +", x%anno%sett, "giorni")

}
