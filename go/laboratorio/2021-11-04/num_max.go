package main

import "fmt"

func main() {
	var x, max, c int
	for i := 0; i < 10; i++ {
		fmt.Print("numero ", i+1, ": ")
		fmt.Scan(&x)
		if x == max {
			c++
		} else if x > max {
			max = x
			c = 1
		}

	}
	fmt.Println(max, "è stato inserito", c, "volte")

}
