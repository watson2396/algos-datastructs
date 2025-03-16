package datastructs

type DNode struct {
	Name string
	Next *DNode
	Prev *DNode
}

type DList struct {
	Size int
	Head *DNode
	Tail *DNode
}

func DListInit() DList {
	return DList{
		Size: 0,
		Head: nil,
		Tail: nil,
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
		d.Tail = n
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
		d.Tail = n
	} else {
		n.Next = element.Next
		n.Prev = element

		if element.Next == nil {
			d.Tail = n
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
		d.Tail = n
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
			d.Tail = element.Prev
		} else {
			element.Next.Prev = element.Prev
		}
	}

	d.Size--
}
