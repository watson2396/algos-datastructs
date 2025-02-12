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
