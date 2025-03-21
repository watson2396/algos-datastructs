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

func (s *Set) SetPrint() {
	s.ListPrint()
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

	for member = s.Head; member != nil; member = s.ListNext(*member) {
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
func (s *Set) SetUnion(s2 *Set) Set {

	sn := Set{
		List{
			Size: 0,
			Head: nil,
		},
	}

	for member := s.Head; member != nil; member = s.ListNext(*member) {
		sn.ListInsNext(member, member.Name)
	}

	for member := s2.Head; member != nil; member = s2.ListNext(*member) {
		if sn.SetIsMember(member) {
			continue
		} else {
			sn.ListInsNext(member, member.Name)
		}

	}

	return sn
}

// Set intersection
func (s *Set) SetIntersection(s2 *Set) Set {
	sn := Set{
		List{
			Size: 0,
			Head: nil,
		},
	}

	for member := s2.Head; member != nil; member = s2.ListNext(*member) {
		if s.SetIsMember(member) {
			sn.ListInsNext(member, member.Name)
		}
	}

	return sn
}

// Set difference
func (s *Set) SetDifference(s2 *Set) Set {
	sn := Set{
		List{
			Size: 0,
			Head: nil,
		},
	}

	for member := s2.Head; member != nil; member = s2.ListNext(*member) {
		if !s.SetIsMember(member) {
			sn.ListInsNext(member, member.Name)
		}
	}

	return sn
}

// Set is member
func (s *Set) SetIsMember(n *Node) bool {
	for member := s.Head; member != nil; member = s.ListNext(*member) {
		if member.Name == n.Name {
			return true
		}
	}

	return false
}

// Set is subset

func (s *Set) SetIsSubset(s2 *Set) bool {
	if s.Size < s2.Size {
		return false
	}

	for member := s2.Head; member != nil; member = s2.ListNext(*member) {
		if !s.SetIsMember(member) {
			return false
		}
	}

	return true
}

// Set is equal
func (s *Set) SetIsEqual(s2 *Set) bool {

	if s.Size != s2.Size {
		return false
	}

	return s.SetIsSubset(s2)
}
