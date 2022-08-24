package main

import "fmt"

type Fifo interface {
	Push(value func())
	Pop() func()
}

type FifoImp struct {
	Fifo  []func()
	Index int
}

func fifoImpl() Fifo {
	return &FifoImp{
		Fifo:  make([]func(), 0),
		Index: 0,
	}
}

func (a *FifoImp) Push(value func()) {
	a.Fifo = append(a.Fifo, value)
	a.Index++
}

func (a *FifoImp) Pop() func() {
	elementRemoved := a.Fifo[len(a.Fifo)-1]
	a.Fifo = append(a.Fifo[:len(a.Fifo)-1])

	return elementRemoved
}

func main() {
	fifoImpl := fifoImpl()

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

	elementRemoved := fifoImpl.Pop()
	elementRemoved()
	elementRemoved = fifoImpl.Pop()
	elementRemoved()
	elementRemoved = fifoImpl.Pop()
	elementRemoved()
}
