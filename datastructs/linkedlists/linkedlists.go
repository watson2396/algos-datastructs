package linkedlists

import "fmt"

type SList struct {
	size int
	head *Node
}

type Node struct {
	next *Node
	name string
}

func SListInit() SList {
	return SList{
		size: 0,
		head: nil,
	}
}

func (s *SList) SListPrepend(name string) {
	n := &Node{
		next: nil,
		name: name,
	}

	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
	s.size++
}

func (s *SList) SListAppend(name string) {
	n := &Node{
		next: nil,
		name: name,
	}

	if s.head == nil {
		s.head = n
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}

	s.size++
}

func (s *SList) SListRemNext(prevNode *Node) {
	// pass in the pointer of the node BEFORE the
	// node you want to remove so you don't
	// have to find the prev one in the method
	// a -> b -> c
	// pass in a to remove b
	// a -> c
	// if pass in NULL, remove a
	// b -> c

	if s.size == 0 {
		return
	}

	// remove Head
	if prevNode == nil {
		var old_element = s.head
		s.head = old_element.next
	} else {
		if prevNode.next == nil {
			return
		}
		var old_element = prevNode.next
		prevNode.next = old_element.next
	}

	s.size--
}

func (l *SList) SListPrint() {
	node := l.head
	for node != nil {
		nm := node.name
		fmt.Println("Name: ", nm)
		node = node.next
	}
}

type DNode struct {
	name string
	next *DNode
	prev *DNode
}

type dList struct {
	size int
	head *DNode
	tail *DNode
}

func DListInit() dList {
	return dList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (d *dList) DListPrepend(name string) {
	n := &DNode{
		name: name,
		next: nil,
		prev: nil,
	}

	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		n.next = d.head
		d.head.prev = n
		d.head = n
	}
}

func (d *dList) DListInsNext(name string, element *DNode) {
	n := &DNode{
		name: name,
		next: nil,
		prev: nil,
	}

	if d.size == 0 {
		d.head = n
		d.tail = n
	} else {
		n.next = element.next
		n.prev = element

		if element.next == nil {
			d.tail = n
		} else {
			element.next.prev = n
		}

		element.next = n
	}

	d.size++
}

func (d *dList) DListInsPrev(name string, element *DNode) {
	n := &DNode{
		name: name,
		next: nil,
		prev: nil,
	}

	if d.size == 0 {
		d.head = n
		d.tail = n
	} else {
		n.next = element
		n.prev = element.prev

		if element.prev == nil {
			d.head = n
		} else {
			element.prev.next = n
		}
	}

	d.size++
}

func (d *dList) DListRemove(element *DNode) {
	if d.size == 0 {
		return
	}

	if d.head == element {
		d.head = element.next
		element.next.prev = nil
	} else {
		element.prev.next = element.next

		if element.next == nil {
			d.tail = element.prev
		} else {
			element.next.prev = element.prev
		}
	}

	d.size--
}
