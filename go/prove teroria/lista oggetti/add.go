package main

func (first *Node) addFirst(x string) *Node {
	var nuovo *Node

	nuovo = new(Node)
	nuovo.value = x
	nuovo.next = first

	return nuovo
}

func (first *Node) addInOrder(x string) *Node {
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
		return first.addFirst(x)
	}
	prev.next = nuovo
	nuovo.next = curr
	return first

}

func (first *Node) addLast(x string) *Node {
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
