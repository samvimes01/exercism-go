package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type List struct {
	head *Node
	tail *Node
}
type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

var ErrEmptyList = errors.New("empty list")

func NewList(elements ...interface{}) *List {
	list := &List{}
	for _, v := range elements {
		list.Push(v)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	currentHead := l.head
	newHead := &Node{Value: v, next: currentHead}
	if currentHead != nil {
		currentHead.prev = newHead
	}
	l.head = newHead
	if l.tail == nil {
		l.tail = newHead
	}
}

func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v, prev: l.tail}
	if l.tail != nil {
		l.tail.next = newNode
	}
	l.tail = newNode
	if l.head == nil {
		l.head = newNode
	}
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}
	currentHead := l.head
	if l.head != nil {
		l.head = l.head.next
	}
	if l.head != nil {
		l.head.prev = nil
	}
	if l.head == nil {
		l.tail = nil
	}
	return currentHead.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		return nil, ErrEmptyList
	}
	currentTail := l.tail
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}
	if l.tail == nil {
		l.head = nil
	}
	return currentTail.Value, nil
}

func (l *List) Reverse() {
	var prev *Node = nil
	var current *Node = l.tail

	l.tail = l.head
	l.head = current

	for current != nil {
		next := current.prev
		current.next = next
		current.prev = prev
		prev = current
		current = next
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
