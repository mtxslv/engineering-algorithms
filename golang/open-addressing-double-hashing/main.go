package main
import (
	"fmt"
	"openaddressing/openaddressing"
)
func main() {
	hashTable := &openaddressing.HashTableOpenAddressingDoubleHashing{}

	hashTable.HashInsert("apple", 100)
	hashTable.HashInsert("banana", 200)
	hashTable.HashInsert("orange", 300)

	fmt.Println("Displaying Hash Table:")
	hashTable.Display()

	item, found := hashTable.Get("banana")
	if found {
		fmt.Printf("Item found:%+v\n", *item)
	} else {
		fmt.Println("'banana' not found")
	}
}
