package stack

import "testing"
func TestPush(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	if s.length != 5 {
		t.Errorf("Expected length of stack to be 5, got %d", s.length)
	}
	if s.top.value != 5 {
		t.Errorf("Expected top of stack to be 5, got %d", s.top.value)
	}
}

func TestPop(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	if s.length != 0 {
		t.Errorf("Expected length of stack to be 0, got %d", s.length)
	}
	if s.top != nil {
		t.Errorf("Expected top of stack to be nil, got %d", s.top.value)
	}
}

func TestIsEmpty(t *testing.T) {
	s := &Stack{}
	if !s.IsEmpty() {
		t.Error("Expected stack to be empty")
	}
}

func TestPeek(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	if s.Peek() != 5 {
		t.Errorf("Expected Peek to return 5, got %d", s.Peek())
	}
}