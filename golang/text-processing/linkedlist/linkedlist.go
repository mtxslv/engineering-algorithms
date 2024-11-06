// linkedlist/linkedlist.go
package linkedlist

import "errors"

type Node[T comparable] struct {
	value *T      // Changed to a pointer
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
	newNode := &Node[T]{value: &v}  // Store a pointer to `v`

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

	if *ll.head.value == v {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}

		ll.len--
		return nil
	}

	current := ll.head
	for current.next != nil {
		if *current.next.value == v {
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

func (n *Node[T]) Value() *T { // Updated to return a pointer
	return n.value
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// toArray traverses the linked list and returns a slice of all values in the list.
func (ll *LinkedList[T]) ToArray() []T {
	// Initialize an empty slice to store the values
	values := make([]T, 0, ll.len)

	// Traverse the linked list
	current := ll.head
	for current != nil {
		values = append(values, *current.value)
		current = current.next
	}

	return values
}