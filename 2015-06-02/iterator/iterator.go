package main

import (
	"fmt"
)

type Iterator interface {
	First() bool
	Next() bool
}

/// -----------------------------
type IterArray struct {
	data  []interface{}
	index int
}

func NewIteratorArray() IterArray {
	out := IterArray{}
	out.index = 0

	return out
}

func (a *IterArray) First() bool {
	a.index = 0

	if len(a.data) > 0 {
		return true
	} else {
		return false
	}
}

func (a *IterArray) Next() bool {
	if a.index < len(a.data) {
		a.index++
		return true
	} else {
		return false
	}
}

func (a *IterArray) GetData(iter Iterator) interface{} {
	return a.data[a.index]
}

func walkIter(a Iterator) {
	if a.First() {
		for nextOne := true; nextOne == true; nextOne = a.Next() {
			fmt.Printf("Current object is: %v\n", a)
		}
	}
}

func main() {
	myArray := NewIteratorArray()

	myArray.data = append(myArray.data, "one")
	myArray.data = append(myArray.data, "two")
	myArray.data = append(myArray.data, "three")

	walkIter(&myArray)

	return
}
