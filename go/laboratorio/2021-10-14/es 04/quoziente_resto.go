package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("dividendo: ")
	fmt.Scan(&a)
	fmt.Print("divisore: ")
	fmt.Scan(&b)
	fmt.Println("quoziente =", a/b)
	fmt.Println("resto =", a%b)

}
