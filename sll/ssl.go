package sll

import "errors"

type SingleLinkedList[T any] struct {
	Head, Tail *Element[T]
}

type Element[T any] struct {
	Next  *Element[T]
	Value T
}

// type SingleLinkedListOperation[T any] interface {
// 	Push(v T)
// 	Unshift(v T)
// 	Pop(v T) error
// 	Shift(v T) error
// 	Find(index int) (*Element[T], error)
// 	Insert(v T, index int) error
// 	Remove(index int) error
// 	Size() int
// 	GetAll() []Element[T]
// }

func (l *SingleLinkedList[T]) Push(v T) {
	if l.Tail == nil {
		l.Head = &Element[T]{Value: v}
		l.Tail = l.Head
	} else {
		l.Tail.Next = &Element[T]{Value: v}
		l.Tail = l.Tail.Next
	}
}

func (l *SingleLinkedList[T]) Unshift(v T) {
	if l.Tail == nil {
		l.Push(v)
		return
	}

	e := &Element[T]{Value: v}
	e.Next = l.Head
	l.Head = e
}

func (l *SingleLinkedList[T]) Pop() error {
	size := l.Size()
	if size == 1 {
		l = &SingleLinkedList[T]{}
		return nil
	}

	node, err := l.Find(size - 2)
	if err != nil {
		return err
	}
	node.Next = nil
	l.Tail = node
	return nil
}

func (l *SingleLinkedList[T]) Shift() error {
	size := l.Size()
	if size == 1 {
		l = &SingleLinkedList[T]{}
		return nil
	}

	if l.Head == nil {
		return errors.New("list is <nil>")
	}

	l.Head = l.Head.Next
	return nil
}

func (l *SingleLinkedList[T]) Insert(v T, index int) error {
	if index == 0 {
		e := &Element[T]{Value: v}
		e.Next = l.Head
		l.Head = e
		return nil
	}

	if index == l.Size() {
		l.Push(v)
		return nil
	}

	node, err := l.Find(index - 1)
	if err != nil {
		return err
	}
	e := &Element[T]{Value: v}
	e.Next = node.Next
	node.Next = e
	return nil
}

func (l *SingleLinkedList[T]) Remove(i int) error {
	if i == 0 {
		return l.Pop()
	}

	if i == l.Size()-1 {
		node, _ := l.Find(i - 1)
		node.Next = nil
		l.Tail = node
		return nil
	}

	node, err := l.Find(i - 1)
	if err != nil {
		return err
	}

	if node.Next.Next != nil {
		node.Next = node.Next.Next
	} else {
		node.Next = nil
	}

	return nil
}

func (l *SingleLinkedList[T]) Find(index int) (*Element[T], error) {
	size := l.Size()
	if index < 0 || index >= size {
		return nil, errors.New("index out of range")
	}

	e := l.Head
	for i := 0; i < index; i++ {
		e = e.Next
	}
	return e, nil
}

func (l *SingleLinkedList[T]) Size() int {
	size := 0
	for e := l.Head; e != nil; e = e.Next {
		size += 1
	}
	return size
}

func (l *SingleLinkedList[T]) GetAll() []T {
	var elems []T
	for e := l.Head; e != nil; e = e.Next {
		elems = append(elems, e.Value)
	}
	return elems
}
