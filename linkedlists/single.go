package linkedlists

import "fmt"

type sList struct {
	size int
	head *Node
}

type Node struct {
	next *Node
	name string
}

func Init() sList {
	return sList{
		size: 0,
		head: nil,
	}
}

func (s *sList) SListPrepend(name string) {
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

func (s *sList) SListAppend(name string) {
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

func (s *sList) SListRemNext(prevNode *Node) {
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

func (l *sList) SListPrint() {
	node := l.head
	for node != nil {
		nm := node.name
		fmt.Println("Name: ", nm)
		node = node.next
	}
}
