package datastructs

// last-in, first-out, LIFO

// Struct composition
type Stack struct{ List }

// Pop
func (s *Stack) StackPop(name string) {
	s.List.ListRemNext(nil)
}

// Push
func (s *Stack) StackPush(name string) {
	s.ListPrepend(name)
}

// Peek
func (s *Stack) StackPeek() string {
	if s.Head == nil {
		return ""
	} else {
		return s.Head.Next.Name
	}
}
