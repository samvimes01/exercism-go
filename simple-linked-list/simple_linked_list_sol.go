package linkedlist2

import "errors"

type Element2 struct {
	data int
	next *Element2
}

type List2 struct {
	head *Element2
	size int
}

func NewList2(numbers []int) *List2 {
	output := &List2{}

	for _, number := range numbers {
		output.Push(number)
	}

	return output
}

func (l *List2) Size() int {
	return l.size
}

func (l *List2) Push(element int) {
	l.head = &Element2{element, l.head}
	l.size++
}

func (l *List2) Pop() (int, error) {
	if l.size < 1 {
		return 0, errors.New("no elements")
	}
	deadHead := l.head
	l.head = deadHead.next
	deadHead.next = nil
	l.size--
	return deadHead.data, nil
}

func (l *List2) Array() []int {
	output := make([]int, l.size)
	for i, head := l.size-1, l.head; i > -1; i, head = i-1, head.next {
		output[i] = head.data
	}
	return output
}

func (l *List2) Reverse() *List2 {
	output := &List2{}
	for head := l.head; head != nil; head = head.next {
		output.Push(head.data)
	}
	return output
}