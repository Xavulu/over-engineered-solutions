//This is honestly pretty useless i just thought it would be funny to use a linked list for fizzbuzz
// ¯\_(ツ)_/¯

package main

import (
	"fmt"
)

type node struct {
	value int
	fizzy string
	next  *node
}

func makeList(add *node, list *node) *node {

	if list == nil { //just check if list is empty

		return list

	}
	for i := list; i != nil; i = i.next {

		if i.next == nil {

			i.next = add
			return list

		}

	}

	return list

}

func printList(list *node) {
	for i := list; i != nil; i = i.next {
		if i != nil && i.next != nil {
			fmt.Println("(", i.value, ",", i.fizzy, ")---> ")
		}
		if i.next == nil {
			fmt.Println("(", i.value, ",", i.fizzy, ")---> nil ")
		}
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func divideNconquer(list *node, front *node, back *node) {

//}

func merge(a *node, b *node) *node {
	res := &node{0, "neither", nil}
	current := res
	for a != nil && b != nil {
		if a.value < b.value {
			current.next = a
			a = a.next
		} else {
			current.next = b
			b = b.next
		}
		current = current.next
	}
	if a != nil {
		current.next = a
	}
	if b != nil {
		current.next = b
	}

	return res.next

}

func sort(list *node) *node { //this part uses recursion and the divide function to merge the seperated nodes
	//just a merge sort to make sure the linked list is in order, this isn't really necessary either
	//base case is checking if list is just one node, we dont need to sort anymore then
	if list == nil || list.next == nil {
		return list
	}

	slow, fast := list, list.next //fast should be ahead
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next //its 2 ahead
	}
	tempHead := slow.next
	slow.next = nil

	return merge(sort(list), sort(tempHead))

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func fizzBuzz(list *node) *node {
	//check if value stored in node is /3 , /5 or both
	//adds fizz/buzz/fizzbuzz to the respective node's value field (fizzy) depending on which one lol
	if list == nil {

		return list

	}
	for i := list; i != nil; i = i.next {

		if i.value%3 == 0 {

			i.fizzy = "Fizz"

		}
		if i.value%5 == 0 {

			i.fizzy = "Buzz"

		}
		if i.value%15 == 0 { //ifits modulo 5 and modulo 3 equal 0 so does modulo 15

			i.fizzy = "FizzBuzz"

		}

	}

	return list

}

func main() {

	node1 := &node{15, "neither", nil} //just initializing 20 nodes
	node2 := &node{1, "neither", nil}
	node3 := &node{9, "neither", nil}
	node4 := &node{11, "neither", nil} //while the original fizzbuzz wants the list from 0 to 100
	node5 := &node{10, "neither", nil} //I just went to 20 because of how unweildly this already is
	node6 := &node{16, "neither", nil}
	node7 := &node{20, "neither", nil}
	node8 := &node{19, "neither", nil}
	node9 := &node{17, "neither", nil}
	node10 := &node{7, "neither", nil}
	node11 := &node{6, "neither", nil}
	node12 := &node{8, "neither", nil}
	node13 := &node{5, "neither", nil}
	node14 := &node{2, "neither", nil}
	node15 := &node{18, "neither", nil}
	node16 := &node{14, "neither", nil}
	node17 := &node{12, "neither", nil}
	node18 := &node{13, "neither", nil}
	node19 := &node{4, "neither", nil}
	node20 := &node{3, "neither", nil}

	head := node1 //setting node head

	head = makeList(node2, head)
	head = makeList(node3, head)
	head = makeList(node4, head)
	head = makeList(node5, head)
	head = makeList(node6, head)
	head = makeList(node7, head)
	head = makeList(node8, head)
	head = makeList(node9, head)
	head = makeList(node10, head)
	head = makeList(node11, head)
	head = makeList(node12, head)
	head = makeList(node13, head)
	head = makeList(node14, head)
	head = makeList(node15, head)
	head = makeList(node16, head)
	head = makeList(node17, head)
	head = makeList(node18, head)
	head = makeList(node19, head)
	head = makeList(node20, head)

	//printList(head)

	fmt.Println("<UNSORTED><UNSORTED><UNSORTED><UNSORTED><UNSORTED><UNSORTED><UNSORTED><UNSORTED><UNSORTED><UNSORTED>")

	printList(head)

	fmt.Println("<SORTED><SORTED><SORTED><SORTED><SORTED><SORTED><SORTED><SORTED><SORTED><SORTED>")

	head = fizzBuzz(sort(head))

	printList(head)
}
