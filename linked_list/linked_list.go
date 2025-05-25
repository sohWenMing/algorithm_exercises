package linkedlist

import (
	"fmt"
	"strings"
)

// ------------Node Definitions and Methods ------------------
type Node struct {
	val  int
	next *Node
}

func InitNode(val int) *Node {
	return &Node{
		val,
		nil,
	}
}

func (n *Node) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(
		"node val: %d\n", n.val,
	))
	switch n.next != nil {
	case true:
		b.WriteString(fmt.Sprintf(
			"node next val: %d\n", n.next.val,
		))
	case false:
		b.WriteString(fmt.Sprint(
			"node next val: nil\n",
		))
	}
	return b.String()
}

// ------------Linked List Definitions and Methods ------------------

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func InitLinkedList() *LinkedList {
	return &LinkedList{
		nil,
		nil,
		0,
	}
}

// define a method to add a new node to the head of the list
func (l *LinkedList) AddToHead(node *Node) {
	return
}

// define a method to add a new node to the tail of the list
func (l *LinkedList) AddToTail(node *Node) {
	return
}

func (l *LinkedList) GetNodeAtIndex(searchIndex int) (node *Node, isFound bool) {
	node = nil
	isFound = false
	return
}

func (l *LinkedList) GetNodeWithVal(searchVal int) (node *Node, index int, isFound bool) {
	node = nil
	index = 0
	isFound = false

	return
}

func (l *LinkedList) AddNodeAtIndex(node *Node, index int) (err error) {
	return nil
}
func (l *LinkedList) RemoveNodeAtIndex(index int) (err error) {
	return nil
}

func (l *LinkedList) IterNodes() (nodes []*Node, vals []int) {
	return nil, nil
}

func (l *LinkedList) GetStringReprs() []string {
	returnSlice := []string{}
	currNode := l.GetHead()
	for currNode != nil {
		returnSlice = append(returnSlice, currNode.String())
		currNode = currNode.next
	}
	return returnSlice
}

func (l *LinkedList) GetLength() int {
	return 0
}

func (l *LinkedList) SetHead(node *Node) {
	return
}
func (l *LinkedList) SetTail(node *Node) {
	return
}

func (l *LinkedList) GetHead() *Node {
	return nil
}
func (l *LinkedList) GetTail() *Node {
	return nil
}
