package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) addFirst(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) addLast(n *node) {
	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = n

	l.length++
}

func (l linkedList) getFirst() int {
	return l.head.data
}

func (l linkedList) getLast() int {
	lastNode := l.head
	for lastNode.next != nil {
		if lastNode.next == nil {
			return 0
		}
		lastNode = lastNode.next
	}

	return lastNode.data
}

func (l linkedList) printLIstData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteFirstHit(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next

	l.length--
}

func (l *linkedList) deleteAllHits(value int) {
	for {
		if l.head.data == value {
			l.head = l.head.next
			l.length--
		}

		previousToDelete := l.head
		for previousToDelete.next.data != value {
			if previousToDelete.next.next == nil {
				return
			}
			previousToDelete = previousToDelete.next
		}
		previousToDelete.next = previousToDelete.next.next

		l.length--
	}
}

func (l linkedList) search(value int) *node {
	if l.length == 0 {
		return nil
	}

	if l.head.data == value {
		return l.head
	}

	searchEl := l.head
	for searchEl.data != value {
		if searchEl.next == nil {
			return nil
		}
		searchEl = searchEl.next
	}

	return searchEl
}

func (n *node) setValue(value int) {
	n.data = value
}

func (l *linkedList) reverse() {
	current := l.head
	var top *node = nil
	for current != nil {
		temp := current.next
		current.next = top
		top = current
		current = temp
	}

	l.head = top
}

func main() {
	mylist := linkedList{}
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
	mylist.reverse()
	mylist.printLIstData()
	//mylist.printLIstData()

	//fmt.Println(mylist.getFirst())
	//fmt.Println(mylist.getLast())

	node13348 := mylist.search(13348)

	if node13348 != nil {
		node13348.setValue(111)
	}

	//mylist.printLIstData()

	//mylistReverse := linkedList{}
	//node11 := &node{data: 11}
	//node22 := &node{data: 22}
	//mylistReverse.addFirst(node11)
	//mylistReverse.addFirst(node22)

}
