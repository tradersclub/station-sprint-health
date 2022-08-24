package main

import "fmt"

type Fifo interface {
	Push(value func())
	Pop() func()
}

type FifoImp struct {
	Fifo []func()
}

func fifoImpl() Fifo {
	return &FifoImp{
		Fifo: make([]func(), 0),
	}
}

func (a *FifoImp) Push(value func()) {
	a.Fifo = append(a.Fifo, value)
}

func (a *FifoImp) Pop() func() {
	if len(a.Fifo) <= 0 {
		return nil
	}

	elementRemoved := a.Fifo[0]
	a.Fifo = a.Fifo[1:]

	return elementRemoved
}

func main() {
	fifoImpl := fifoImpl()

	elementRemoved := fifoImpl.Pop()
	if elementRemoved != nil {
		elementRemoved()
	}

	fifoImpl.Push(func() {
		fmt.Println("1")
	})
	fifoImpl.Push(func() {
		fmt.Println("2")
	})
	fifoImpl.Push(func() {
		fmt.Println("3")
	})
	fmt.Println(fifoImpl)

	elementRemoved = fifoImpl.Pop()
	elementRemoved()
	elementRemoved = fifoImpl.Pop()
	elementRemoved()
	elementRemoved = fifoImpl.Pop()
	elementRemoved()
}
