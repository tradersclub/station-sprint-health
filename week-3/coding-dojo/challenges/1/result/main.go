package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value func()
	next  *Node
}

type Queue struct {
	root *Node
	tail *Node
}

type IQueue interface {
	Pop() (func(), error)
	Push(newTask func())
}

func newQueue() IQueue {
	return &Queue{}
}

func (q *Queue) Push(newTask func()) {
	newNode := &Node{value: newTask}

	if q.root == nil {
		q.root = newNode
		q.tail = q.root
		return
	}

	q.tail.next = newNode
	q.tail = q.tail.next
}

func (q *Queue) Pop() (func(), error) {

	lastRoot := q.root

	if lastRoot == nil {
		q.tail = q.root
		return nil, errors.New("empty queue")
	}

	q.root = lastRoot.next

	return lastRoot.value, nil
}

func main() {
	queue := newQueue()

	queue.Push(func() {
		fmt.Println("1")
	})
	queue.Push(func() {
		fmt.Println("20")
	})
	queue.Push(func() {
		fmt.Println("10")
	})

	task1, err := queue.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	task1()

	task2, err := queue.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	task2()

	queue.Push(func() {
		fmt.Println("39")
	})

	task3, err := queue.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	task3()

	task4, err := queue.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	task4()

	task5, err := queue.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	task5()
}
