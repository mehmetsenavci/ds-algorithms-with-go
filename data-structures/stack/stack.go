package stack

type node struct {
	value interface{}
	next  *node
}

type Stack struct {
	top *node
	length int
}

func (s *Stack) Push(value interface{}) {
	s.top = &node{value, s.top}
	s.length++
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.length--
	return value
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}