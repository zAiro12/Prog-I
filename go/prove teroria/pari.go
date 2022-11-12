package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)
	if Pari(x) {
		fmt.Println("pari")
	} else {
		fmt.Println("dispari")
	}

}

func Pari(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}
