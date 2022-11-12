package main

import "fmt"

func main() {
	var a, aMax byte
	for i := 0; i < 5; i++ {
		fmt.Scanf("%c", &a)
		fmt.Scan()
		//fmt.Println(i)

		if a > aMax {
			aMax = a
		}
	}
	fmt.Println(string(aMax))

}
