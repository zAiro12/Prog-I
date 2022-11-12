package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	for x != 0 {
		if x%10 == 3 {
			break
		}
		x /= 10
	}

	if x != 0 {
		fmt.Println("3 compare")
	} else {
		fmt.Println("3 non compare")
	}
}
