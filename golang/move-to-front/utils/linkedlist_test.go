package utils

import (
	"testing"
)

// type Item struct {
	// Title  string
	// Author string
// }

// Mock access sequence
var accessSequence = []string{"SongB", "SongD", "SongA", "SongC", "SongE"}

// Initialize the list
func initializeList() *LinkedList {
	ll := New()
	items := []*Item{
		{Title: "SongA", Author: "Author1"},
		{Title: "SongB", Author: "Author2"},
		{Title: "SongC", Author: "Author3"},
		{Title: "SongD", Author: "Author4"},
		{Title: "SongE", Author: "Author5"},
	}
	for _, item := range items {
		ll.Add(item)
	}
	return ll
}

// Simulate optimal Foresee cost
func optimalForeseeCost(accessSequence []string) int {
	// Foresee assumes minimal access cost (always finds the item in the first position).
	return len(accessSequence)
}

// Calculate MTF cost
func mtfCost(ll *LinkedList, sequence []string) int {
	totalCost := 0
	for _, song := range sequence {
		current := ll.Head
		cost := 1
		for current != nil {
			if current.Content.Title == song {
				totalCost += cost
				ll.SearchAndMoveToFront(song)
				break
			}
			cost++
			current = current.Next
		}
	}
	return totalCost
}

func TestMoveToFrontVsForesee(t *testing.T) {
	ll := initializeList()

	// Calculate costs
	mtf := mtfCost(ll, accessSequence)
	foresee := optimalForeseeCost(accessSequence)

	// Assert competitiveness
	t.Logf("MTF Cost: %d, Foresee Cost: %d", mtf, foresee)
	t.Logf("MTF/Foresee Cost Ratio: %.4f", float32(mtf)/float32(foresee))
	if mtf > 4*foresee {
		t.Errorf("MTF is not competitive: got %d, want <= %d", mtf, 4*foresee)
	}
}
