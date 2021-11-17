// The concept:
// 1. loop through linked list & save the numbers in hash table
// 2. If there is same number in hash table, then remove it
// 3. for example, [1, 1, 2, 3, 3] -> [1, 2, 3]

package main

import (
	"self/linked_list/singly_linked_list"
)

// init a linked list
func remove_dups(l *singly_linked_list.List) *singly_linked_list.List {
	l.Display()
	return l
}

func main() {
	test := &singly_linked_list.List{}
	test.Insert(1)
	test.Insert(2)
	test.Insert(2)
	test.Insert(4)
	test.Insert(5)
	test.Insert(6)

	remove_dups(test)
}
