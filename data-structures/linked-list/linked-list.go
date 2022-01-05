package linkedlist

type node struct {
	Value interface{}
	Next *node
}

type LinkedList struct {
	Head *node
	Tail *node
	Size int
}

func New() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Add(value interface{}) {
	newNode := &node{Value: value}
	if ll.Size == 0 {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
	ll.Size++
}

func (ll *LinkedList) Remove(index int) {
	if ll.Size == 0 {
		return
	}

	if index == 0 {
		ll.Head = ll.Head.Next
		ll.Size--
		return
	}

	prevNode := ll.Head
	for i := 0; i < index-1; i++ {
		prevNode = prevNode.Next
	}

	prevNode.Next = prevNode.Next.Next
	ll.Size--
}

func (ll *LinkedList) Get(index int) interface{} {
	if ll.Size == 0 {
		return nil
	}

	if index == 0 {
		return ll.Head.Value
	}

	prevNode := ll.Head
	for i := 0; i < index-1; i++ {
		prevNode = prevNode.Next
	}

	return prevNode.Next.Value
}