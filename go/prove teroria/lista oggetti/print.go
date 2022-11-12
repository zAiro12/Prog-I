package main

import "fmt"

func (first *Node) print() {
	for first != nil {
		fmt.Println(first.value)
		//if first.next != nil {
		first = first.next
		//}
	}
}
