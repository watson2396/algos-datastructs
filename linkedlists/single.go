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

func (s *sList) PrependSingle(name string) {
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
	return
}

func (s *sList) AppendSingle(name string) {
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
	return
}

func (l *sList) PrintSingle() {
	node := l.head
	for node != nil {
		nm := node.name
		fmt.Println("Name: ", nm)
		node = node.next
	}
}
