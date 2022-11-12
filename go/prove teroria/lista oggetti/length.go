package main

func (first *Node) lenght() int {
	c := 0
	for first != nil {
		c++
		first = first.next
	}
	return c
}
