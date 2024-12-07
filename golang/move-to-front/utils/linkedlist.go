package utils

import (
	"fmt"
)

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

func (ll *LinkedList) MoveToFront(node *Node) {
	// Move a given node to the front of the list.
	rightNode := node
	for rightNode.Previous != nil { // While there is a previous node
		leftNode := rightNode.Previous
		err := ll.SwapNodes(leftNode, rightNode)
		if err != nil {
			panic(err)
		}
	}
}


func (ll *LinkedList) SwapNodes(leftNode, rightNode *Node) error {
/*
Suppose the front of the list (first node) is on the left:

					A == B == C == D == E

A is the front of the list.
So let's always refer to the nodes swapping position as leftNode and rightNode such that, before swapping:


					alphaNode == leftNode == rightNode == omegaNode

and after swapping

					alphaNode == rightNode == leftNode == omegaNode

The operations then must be:
	rightNode.Previous = &alphaNode
	rightNode.Next = &leftNode
	leftNode.Previous = &rightNode
	leftNode.Next = &omegaNode
	alphaNode.Next = &rightNode
	omegaNode.Previous = &leftNode
*/

	if leftNode == nil || rightNode == nil || leftNode == rightNode {
		return fmt.Errorf("Error occurred: %w", ErrNodesInvalid)
	}
	// Ensure leftNode comes right before rightNode 
	if leftNode.Next != rightNode {
		return fmt.Errorf("Error occurred: %w",ErrNodesOrderInvalid)
	}

	alphaNode := leftNode.Previous
	omegaNode := rightNode.Next

	// Update rightNode's pointers
	rightNode.Previous = alphaNode
	rightNode.Next = leftNode

	// Update leftNode's pointers
	leftNode.Previous = rightNode
	leftNode.Next = omegaNode

	if alphaNode != nil { // alphaNode exists
		alphaNode.Next = rightNode
	} else { // alphaNode does not exist. 
		ll.Head = rightNode // Now rightNode is the Head
	}

	if omegaNode != nil { // omegaNode exists
		omegaNode.Previous = leftNode
	} else { // omegaNode does not exist
		ll.Tail = leftNode // Now leftNode is the Tail
	}

	return nil

}
