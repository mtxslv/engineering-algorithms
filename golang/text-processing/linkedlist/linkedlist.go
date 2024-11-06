// linkedlist/linkedlist.go
package linkedlist

import "errors"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Add(v T) {
	newNode := &Node[T]{value: v}

	if ll.len == 0 {
		ll.head = newNode
	} else {
		ll.tail.next = newNode
	}

	ll.tail = newNode
	ll.len++
}

func (ll *LinkedList[T]) Remove(v T) error {
	if ll.len == 0 {
		return ErrEmptyList
	}

	if ll.head.value == v {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}

		ll.len--
		return nil
	}

	current := ll.head
	for current.next != nil {
		if current.next.value == v {
			current.next = current.next.next
			if current.next == nil {
				ll.tail = current
			}

			ll.len--
			return nil
		}

		current = current.next
	}

	return ErrItemNotFound
}

var (
	ErrEmptyList    = errors.New("empty list")
	ErrItemNotFound = errors.New("item not found")
)

func (ll *LinkedList[T]) Length() int {
	return ll.len
}

func (ll *LinkedList[T]) Head() *Node[T] {
	return ll.head
}

// linkedlist/linkedlist.go
func (n *Node[T]) Value() T {
	return n.value
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

