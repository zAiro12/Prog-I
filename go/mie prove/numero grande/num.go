package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	if a > b && a > c && a > d && a > e {
		fmt.Println(a, "è il numero più grande")
	} else if b > a && b > c && b > d && b > e {
		fmt.Println(b, "è il numero più grande")
	} else if c > b && c > a && c > d && c > e {
		fmt.Println(c, "è il numero più grande")
	} else if d > b && d > c && d > a && d > e {
		fmt.Println(d, "è il numero più grande")
	} else {
		fmt.Println(e, "è il numero più grande")
	}
}
