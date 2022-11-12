package main

import "fmt"

func main() {
	var first *Node

	first = addInOrder(first, "ciao")
	first = addInOrder(first, "banana")
	first = addInOrder(first, "zorro")
	first = addInOrder(first, "anaconda")

	fmt.Println(length(first))

	print(first)
}
