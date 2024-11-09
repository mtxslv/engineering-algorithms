package main
import (
	"fmt"
	"openaddressing/openaddressing"
)
func main() {
	hashTable := &openaddressing.HashTableOpenAddressing{}

	hashTable.Insert("apple", 100)
	hashTable.Insert("banana", 200)
	hashTable.Insert("orange", 300)

	fmt.Println("Displaying Hash Table:")
	hashTable.Display()

	value, found := hashTable.Search("banana")
	if found {
		fmt.Printf("Found 'banana' with value %d\n", value)
	} else {
		fmt.Println("'banana' not found")
	}

	hashTable.Delete("banana")
	fmt.Println("Displaying Hash Table after deleting 'banana':")
	hashTable.Display()
}
