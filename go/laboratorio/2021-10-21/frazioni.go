package main

import "fmt"

func main() {
	var num1, den1, num2, den2 int
	fmt.Print("num e den fraz 1: ")
	fmt.Scan(&num1, &den1)
	fmt.Print("num e den fraz 2: ")
	fmt.Scan(&num2, &den2)
	if (num1%num2 == 0 && den1%den2 == 0) || (num2%num1 == 0 && den2%den1 == 0) {
		fmt.Println("equivalenti")
	} else {
		fmt.Println("non equivalenti")
	}
}
