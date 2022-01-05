package linkedlist

import "testing"

func TestFirstAdd(t *testing.T) {
	ll := New()
	ll.Add(1)
	if ll.Size != 1 {
		t.Errorf("Expected size to be 1, got %d", ll.Size)
	}
	if ll.Head.Value != 1 {
		t.Errorf("Expected head value to be 1, got %d", ll.Head.Value)
	}
	if ll.Tail.Value != 1 {
		t.Errorf("Expected tail value to be 1, got %d", ll.Tail.Value)
	}
}

func TestAdd(t *testing.T) {
	ll := New()
	ll.Add(1)
	ll.Add(2)
	if ll.Size != 2 {
		t.Errorf("Expected size to be 2, got %d", ll.Size)
	}
	if ll.Head.Value != 1 {
		t.Errorf("Expected head value to be 1, got %d", ll.Head.Value)
	}
	if ll.Tail.Value != 2 {
		t.Errorf("Expected tail value to be 2, got %d", ll.Tail.Value)
	}
}

func TestRemoveFirstIndex(t *testing.T) {
	ll := New()
	ll.Add(1)
	ll.Add(2)
	ll.Remove(0)
	if ll.Size != 1 {
		t.Errorf("Expected size to be 1, got %d", ll.Size)
	}
	if ll.Head.Value != 2 {
		t.Errorf("Expected head value to be 2, got %d", ll.Head.Value)
	}
	if ll.Tail.Value != 2 {
		t.Errorf("Expected tail value to be 2, got %d", ll.Tail.Value)
	}
}

func TestRemove(t *testing.T) {
	ll := New()
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Remove(1)
	if ll.Size != 2 {
		t.Errorf("Expected size to be 2, got %d", ll.Size)
	}
	if ll.Head.Value != 1 {
		t.Errorf("Expected head value to be 1, got %d", ll.Head.Value)
	}
	if ll.Tail.Value != 3 {
		t.Errorf("Expected tail value to be 3, got %d", ll.Tail.Value)
	}
}

