package main

func addFirst(first *Node, x string) *Node {
	var nuovo *Node

	nuovo = new(Node)
	nuovo.value = x
	nuovo.next = first

	return nuovo
}

func addInOrder(first *Node, x string) *Node {
	var nuovo *Node

	nuovo = new(Node)
	nuovo.value = x

	var prev, curr *Node
	prev = nil
	curr = first
	for curr != nil && curr.value < x {
		prev = curr
		curr = curr.next
	}
	if prev == nil {
		return addFirst(first, x)
	}
	prev.next = nuovo
	nuovo.next = curr
	return first

}

func addLast(first *Node, x string) *Node {
	var nuovo *Node
	nuovo = new(Node)
	nuovo.value = x
	nuovo.next = nil
	if first == nil {
		return nuovo
	}
	c := first
	for c.next != nil {
		c = c.next
	}
	c.next = nuovo
	return first
}
