package main

import "fmt"

func print(first *Node) {
	for first != nil {
		fmt.Println(first.value)
		//if first.next != nil {
		first = first.next
		//}
	}
}
