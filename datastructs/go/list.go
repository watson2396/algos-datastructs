package datastructs

import "fmt"

type List struct {
	Size int
	Head *Node
	Tail *Node
}

type Node struct {
	Next *Node
	Name string
}

func ListInit() List {
	return List{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
}

func (s *List) ListPrepend(Name string) {
	n := &Node{
		Next: nil,
		Name: Name,
	}

	if s.Head == nil {
		s.Head = n
		s.Tail = n
	} else {
		n.Next = s.Head
		s.Head = n
	}
	s.Size++
}

func (s *List) ListAppend(Name string) {
	n := &Node{
		Next: nil,
		Name: Name,
	}

	if s.Head == nil {
		s.Head = n
		s.Tail = n
	} else {
		current := s.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = n
	}

	s.Size++
}

// Insert after the given node pointer
func (s *List) ListInsNext(a *Node, name string) {
	n := &Node{
		Next: nil,
		Name: name,
	}

	// if position to insert after is null
	// assume it is an empty list
	if a == nil {
		s.Head = n
		s.Tail = n
	} else {

		if a.Next == nil {
			s.Tail = n
		}
		n.Next = a.Next
		a.Next = n
	}

	s.Size++
}

func (s *List) ListRemNext(PrevNode *Node) string {
	// pass in the pointer of the node BEFORE the
	// node you want to remove so you don't
	// have to find the Prev one in the method
	// a -> b -> c
	// pass in a to remove b
	// a -> c
	// if pass in NULL, remove a
	// b -> c

	if s.Size == 0 {
		return ""
	}

	// remove Head
	if PrevNode == nil {
		var old_element = s.Head
		s.Head = old_element.Next
		s.Size--
		return old_element.Name
	} else {
		if PrevNode.Next == nil {
			return ""
		}
		var old_element = PrevNode.Next
		PrevNode.Next = old_element.Next
		s.Size--
		return old_element.Name
	}

}

func (l *List) ListPrint() {
	for node := l.Head; node != nil; node = node.Next {
		fmt.Println("Name:", node.Name)
	}
}

func (l *List) ListNext(n Node) *Node {
	return n.Next
}
