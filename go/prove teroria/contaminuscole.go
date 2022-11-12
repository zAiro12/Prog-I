package main

import "fmt"

func main() {
	var s string
	fmt.Print("Inserisci una stringa: ")
	fmt.Scan(&s)
	voc, con := conta(s)
	fmt.Println("Vocali minuscole:", voc)
	fmt.Println("Consonanti minuscole:", con)
}

func conta(s string) (int, int) {
	cv := 0
	cc := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			cv++

		case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
			cc++
		}
	}
	return cv, cc
}

//OPPURE

// func conta(s string) (cv int, cc int) {
// 	cv = 0
// 	cc = 0
// 	for i := 0; i < len(s); i++ {
// 		switch s[i] {
// 		case 'a', 'e', 'i', 'o', 'u':
// 			cv++

// 		case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
// 			cc++
// 		}
// 	}
// 	return
// }
