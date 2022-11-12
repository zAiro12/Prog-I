package main

import "fmt"

func main() {
	var first *Node

	first.addInOrder("ciao")
	first.addInOrder("banana")
	first.addInOrder("zorro")
	first.addInOrder("anaconda")

	fmt.Println(first.lenght())

	first.print()
}
