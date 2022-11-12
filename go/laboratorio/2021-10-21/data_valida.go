package main

import "fmt"

func main() {
	var gg, mm, aa int
	fmt.Print("gg mm aa? ")
	fmt.Scan(&gg, &mm, &aa)
	if !veraabis(aa) {
		vermm(gg, mm)
	} else {
		vergg(gg, 29)
	}

}

func vermm(gg int, mm int) {
	if mm == 1 || mm == 3 || mm == 5 || mm == 7 || mm == 8 || mm == 10 || mm == 12 {
		vergg(gg, 31)
	} else if mm == 4 || mm == 6 || mm == 9 || mm == 11 {
		vergg(gg, 30)
	} else if mm == 2 {
		vergg(gg, 28)
	} else {
		fmt.Println("data non valida")
	}
}

func vergg(gg int, k int) {
	if gg >= 0 && gg <= k {
		fmt.Println("data valida")
	} else {
		fmt.Println("data non valida")
	}
}

func veraabis(aa int) bool {
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
