package linkedlists

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
		d.size++
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
