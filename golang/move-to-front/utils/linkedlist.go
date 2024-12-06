package utils

// import "fmt"

type Node struct{
	Previous *Node
	Content *Item
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	len int
}

func New() *LinkedList {
	return &LinkedList{}
}

// Add appends an Item to the end of the linked list
func (ll *LinkedList) Add(item *Item) {
	// Create a new node with the given item.
	newNode := &Node{
		Content: item,
		Next: nil, // it will be the last node, so no next node
	}

	if ll.len == 0 {
		// List empty. Then newNode is both Head and Tail
		ll.Head = newNode
		ll.Tail = newNode

	} else {
		// List is not empty, append the node at the end

		// newNode's previous node is current Tail
		newNode.Previous = ll.Tail 
		// current Tail's next node is newNode
		ll.Tail.Next = newNode 
		// newNode is the new Tails
		ll.Tail = newNode
	}
	ll.len++
}

func (ll *LinkedList) Search(title string) *Item {
	// Traverse and print the linked list.
	current := ll.Head
	for current != nil {
		// Ensure current.Content is not nil before dereferencing.
		if current.Content != nil {
			content := *current.Content 
			if content.Title == title {
				return current.Content
			}
		}
		// Move to the next node.
		current = current.Next
	}
	return nil
}