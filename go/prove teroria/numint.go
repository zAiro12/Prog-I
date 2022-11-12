package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	c := 0
	if x == 0 {
		fmt.Println("1")
	} else {
		for x > 0 {
			x /= 10
			c++
		}
		fmt.Println(c)
	}

}
