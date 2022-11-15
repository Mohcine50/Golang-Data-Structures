package linkedlist

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

//Linkedlist travers

func (list *LinkedList) Traverse() {

	for list.Head != nil {
		fmt.Println(list.Head.Value)
		list.Head = list.Head.Next
	}

}

// Linked List node Insertion
/** Insert at the beginnin of the linedlist**/
func (list *LinkedList) InsertInBeginin(value int) {
	node := &Node{Value: value, Next: nil}
	if list.Head == nil {
		list.Head = node
	} else {
		node.Next = list.Head
		list.Head = node
	}
}

/** Insert at the end of the linedlist**/
func (list *LinkedList) InsertInEnd(value int) {
	node := &Node{Value: value, Next: nil}
	if list.Head == nil {
		list.Head = node
	} else {
		var tempNode *Node = list.Head
		for tempNode.Next != nil {
			tempNode = tempNode.Next
		}
		tempNode.Next = node
	}
}

//delete the head node
func (list *LinkedList) RemoveHead() {

	list.Head = list.Head.Next

}

// delete tail node
func (list *LinkedList) RemoveTail() {
	var tempNode *Node = list.Head
	for tempNode.Next.Next != nil {
		tempNode = tempNode.Next
	}
	tempNode.Next = nil

}
