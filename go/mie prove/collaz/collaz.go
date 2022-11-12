package main

import "fmt"

func main() {
	num := 1
	ok := 0
	fmt.Println(num)
	for{
		for i := num; i != 1; {
			i = infi(i)
		}
		ok++
		fmt.Println(ok)
		num++
	}

}

func infi(n int) int{
	if n == 1 {
		return n
	}
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
