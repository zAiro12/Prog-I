package main

import "fmt"

func main() {
	var b byte
	fmt.Print("Inserisci una lettera: ")
	fmt.Scanf("%c", &b)
	fmt.Scanf("")
	fmt.Println(b)
	fmt.Printf("%c %c %c \n", b-1, b, b+1)
	if b <= 'l' && b >= 'a' {
		fmt.Println("A-L")
	} else {
		fmt.Println("altro")
	}
	var parola string
	fmt.Print("Inserisci una parola: ")
	fmt.Scan(&parola)
	for _, i := range parola {
		fmt.Println(string(i))
	}
}
