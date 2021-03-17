package main

import (
	"fmt"
)

type node struct {
	data int
	prev *node
	next *node
}

type doublyLinkedList struct {
	head   *node
	last   *node
	length int
}

func (l *doublyLinkedList) addFirst(n *node) {
	if l.length == 0 {
		l.head = n
		l.last = n
		l.length++
		return
	}
	second := l.head
	l.head = n

	l.head.next = second
	second.prev = l.head
	l.length++
}

func (l *doublyLinkedList) addLast(n *node) {
	temp := l.last
	temp.next = n

	l.last = n
	l.last.prev = temp

	l.length++
}

func (l doublyLinkedList) getFirst() int {
	return l.head.data
}

func (l doublyLinkedList) getLast() int {
	return l.last.data
}

func (l doublyLinkedList) printLIstData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

//func (l *doublyLinkedList) deleteFirstHit(value int) {
//
//}
//
//func (l *doublyLinkedList) deleteAllHits(value int) {
//
//}

//func (l doublyLinkedList) search(value int) *node {
//
//}

//func (n *node) setValue(value int) {
//
//}
//
func (l *doublyLinkedList) reverse() {
	node := l.head
	for node != nil {
		node.next, node.prev = node.prev, node.next
		node = node.prev
	}

	l.head, l.last = l.last, l.head
}

func main() {
	mylist := doublyLinkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 158}
	node3 := &node{data: 148}
	node4 := &node{data: 159}
	node5 := &node{data: 148}
	node6 := &node{data: 1148}
	node7 := &node{data: 13348}
	node8 := &node{data: 1}

	mylist.addFirst(node1)
	mylist.addFirst(node2)
	mylist.addFirst(node3)
	mylist.addFirst(node4)
	mylist.addFirst(node5)
	mylist.addFirst(node6)
	mylist.addFirst(node7)
	mylist.addLast(node8)

	mylist.printLIstData()

	fmt.Println(mylist.getFirst())
	fmt.Println(mylist.getLast())

	mylist.reverse()

	fmt.Println(mylist.getFirst())
	fmt.Println(mylist.getLast())

	mylist.printLIstData()
}
