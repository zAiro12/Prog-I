package main

import "fmt"

func main() {
	var x int
	fmt.Print("voto: ")
	fmt.Scan(&x)
	if x >= 30 {
		fmt.Println("voto non valido")
	}
}
