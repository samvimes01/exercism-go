package linkedlist2

import "errors"

// Define the List and Element types here.
type Element struct {
	Value int
	next  *Element
}
type List struct {
	head *Element
	tail *Element
	size int
}

var ErrEmptyList = errors.New("empty list")

func New(elements []int) *List {
	list := &List{}
	var prev *Element = nil
	for _, element := range elements {
		el := &Element{Value: element}
		if prev == nil {
			list.head = el
		} else {
			prev.next = el
		}
		list.tail = el
		prev = el
		list.size++
	}
	return list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	el := &Element{Value: element}
	if l.size == 0 {
		l.head = el
		l.tail = el
	} else {
		l.tail.next = el
		l.tail = el
	}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, ErrEmptyList
	}
	el := l.head
	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size--
		return el.Value, nil
	}
	for i := 0; i < l.size-1; i++ {
		next := el.next
		if i == l.size-2 {
			el.next = nil
			l.tail = el
		}
		el = next
	}
	l.size--
	return el.Value, nil
}

func (l *List) Array() []int {
	list := make([]int, 0)
	for el := l.head; el != nil; el = el.next {
		list = append(list, el.Value)
	}
	return list
}

func (l *List) Reverse() *List {
	var prev *Element = nil
	var current *Element = l.head

	l.tail = l.head
	l.head = current

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
		l.head = prev
	}
	return l
}
