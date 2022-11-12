package main

func length(first *Node) int {
	c := first
	for c != nil {
		c++
		c = c.next
	}
	return c
}
