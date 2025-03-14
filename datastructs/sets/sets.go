package sets

import (
	"github.com/watson2396/algos-datastructs/datastructs/linkedlists"
)

type Set struct{ linkedlists.List }

func SetInit() Set {
	return Set{
		linkedlists.List{
			Size: 0,
			Head: nil,
		},
	}
}

// Set insert
func (s *Set) SetInsert(name string) bool {
	n := &linkedlists.Node{
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
func (s *Set) SetRemove(n *linkedlists.Node) {
	member := &linkedlists.Node{}
	prev := &linkedlists.Node{}

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

func (s *Set) SetIsMember(n *linkedlists.Node) bool {
	member := &linkedlists.Node{}

	for member = s.Head; member == nil; member = s.ListNext(*member) {
		if member.Name == n.Name {
			return true
		}
	}

	return false
}

// Set is subset

// Set is equal
