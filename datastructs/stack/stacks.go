package stacks

// last-in, first-out, LIFO

import (
	"github.com/watson2396/algos-datastructs/datastructs/linkedlists"
)

// Struct composition
type Stack struct{ linkedlists.SList }

// Pop
func (s *Stack) StackPop(name string) {
	s.SList.SListRemNext(nil)
}

// Push
func (s *Stack) StackPush(name string) {
	s.SListPrepend(name)
}

// Peek
func (s *Stack) StackPeek() string {
	if s.Head == nil {
		return ""
	} else {
		return s.Head.Next.Name
	}
}
