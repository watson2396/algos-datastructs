package datastructs

type Set struct{ List }

func SetInit() Set {
	return Set{
		List{
			Size: 0,
			Head: nil,
		},
	}
}

// Set insert
func (s *Set) SetInsert(name string) bool {
	n := &Node{
		Name: name,
		Next: nil,
	}

	if s.SetIsMember(n) {
		return true
	}

	s.ListInsNext(s.Tail, name)
	return true

}

// Set remove
func (s *Set) SetRemove(n *Node) {
	member := &Node{}
	prev := &Node{}

	for member = s.Head; member == nil; member = s.ListNext(*member) {
		if member.Name == n.Name {
			break
		}
		prev = member
	}

	if member == nil {
		return
	}

	s.ListRemNext(prev)
}

// Set union

// Set intersection

// Set difference

// Set is member

func (s *Set) SetIsMember(n *Node) bool {
	member := &Node{}

	for member = s.Head; member == nil; member = s.ListNext(*member) {
		if member.Name == n.Name {
			return true
		}
	}

	return false
}

// Set is subset

// Set is equal
