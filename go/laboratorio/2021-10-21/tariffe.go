package main

import "fmt"

func main() {
	var eta int
	var stud bool
	fmt.Print("età: ")
	fmt.Scan(&eta)
	fmt.Print("studente? (t/f): ")
	fmt.Scan(&stud)
	if stud && eta >= 18 {
		fmt.Println("ingresso 5 euro")
	} else {
		if eta >= 0 && eta < 9 {
			fmt.Println("ingresso gratis")
		} else if eta >= 9 && eta < 18 {
			fmt.Println("ingresso 5 euro")
		} else if eta >= 18 && eta < 26 {
			fmt.Println("ingresso 7.5 euro")
		} else if eta >= 26 && eta < 65 {
			fmt.Println("ingresso 10 euro")
		} else {
			fmt.Println("ingresso 7.5 euro")
		}
	}
}
