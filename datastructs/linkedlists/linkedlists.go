package linkedlists

import "fmt"

type SList struct {
	Size int
	Head *Node
}

type Node struct {
	Next *Node
	Name string
}

func SListInit() SList {
	return SList{
		Size: 0,
		Head: nil,
	}
}

func (s *SList) SListPrepend(Name string) {
	n := &Node{
		Next: nil,
		Name: Name,
	}

	if s.Head == nil {
		s.Head = n
	} else {
		n.Next = s.Head
		s.Head = n
	}
	s.Size++
}

func (s *SList) SListAppend(Name string) {
	n := &Node{
		Next: nil,
		Name: Name,
	}

	if s.Head == nil {
		s.Head = n
	} else {
		current := s.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = n
	}

	s.Size++
}

func (s *SList) SListRemNext(PrevNode *Node) string {
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

func (l *SList) SListPrint() {
	node := l.Head
	for node != nil {
		nm := node.Name
		fmt.Println("Name: ", nm)
		node = node.Next
	}
}

type DNode struct {
	Name string
	Next *DNode
	Prev *DNode
}

type DList struct {
	Size int
	Head *DNode
	tail *DNode
}

func DListInit() DList {
	return DList{
		Size: 0,
		Head: nil,
		tail: nil,
	}
}

func (d *DList) DListPrepend(Name string) {
	n := &DNode{
		Name: Name,
		Next: nil,
		Prev: nil,
	}

	if d.Head == nil {
		d.Head = n
		d.tail = n
	} else {
		n.Next = d.Head
		d.Head.Prev = n
		d.Head = n
	}
}

func (d *DList) DListInsNext(Name string, element *DNode) {
	n := &DNode{
		Name: Name,
		Next: nil,
		Prev: nil,
	}

	if d.Size == 0 {
		d.Head = n
		d.tail = n
	} else {
		n.Next = element.Next
		n.Prev = element

		if element.Next == nil {
			d.tail = n
		} else {
			element.Next.Prev = n
		}

		element.Next = n
	}

	d.Size++
}

func (d *DList) DListInsPrev(Name string, element *DNode) {
	n := &DNode{
		Name: Name,
		Next: nil,
		Prev: nil,
	}

	if d.Size == 0 {
		d.Head = n
		d.tail = n
	} else {
		n.Next = element
		n.Prev = element.Prev

		if element.Prev == nil {
			d.Head = n
		} else {
			element.Prev.Next = n
		}
	}

	d.Size++
}

func (d *DList) DListRemove(element *DNode) {
	if d.Size == 0 {
		return
	}

	if d.Head == element {
		d.Head = element.Next
		element.Next.Prev = nil
	} else {
		element.Prev.Next = element.Next

		if element.Next == nil {
			d.tail = element.Prev
		} else {
			element.Next.Prev = element.Prev
		}
	}

	d.Size--
}
