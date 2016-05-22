// Generated by: main
// TypeWriter: ChunkedVec
// Directive: +gen on RenderComponent

package components

import (
	"bytes"
	"container/list"
	"fmt"
)

// This is an implementation of https://github.com/mzdravkov/chunked-vector
// The MIT License (MIT)
// Copyright (c) 2016 Mihail Zdravkov (mihail0zdravkov@gmail.com)

type RenderComponentChunkedVec struct {
	List      *list.List
	ChunkSize uint
	Empty     RenderComponent
}

// Creates a new RenderComponentChunkedVec with chunkSize as provided
func NewRenderComponentChunkedVec(chunkSize uint) *RenderComponentChunkedVec {
	if chunkSize == 0 {
		chunkSize = 1024
	}

	return &RenderComponentChunkedVec{
		List:      list.New(),
		ChunkSize: chunkSize,
	}
}

// Adds the element to the ChunkedVec and returns the position it was added to
func (cv *RenderComponentChunkedVec) Add(element RenderComponent) (uint, uint) {
	listIndex := 0
	for e := cv.List.Front(); e != nil; e = e.Next() {
		for index, value := range e.Value.([]RenderComponent) {
			if value == cv.Empty {
				e.Value.([]RenderComponent)[index] = element
				return uint(listIndex), uint(index)
			}
		}

		listIndex++
	}

	slice := make([]RenderComponent, cv.ChunkSize)
	slice[0] = element
	cv.List.PushBack(slice)

	return uint(listIndex), uint(0)
}

// Overwrites the given position to hold the given value
func (cv *RenderComponentChunkedVec) PutAt(element RenderComponent, listIndex, sliceIndex uint) {
	var i uint = 0
	e := cv.List.Front()
	for ; i < listIndex; e = e.Next() {
		i++
	}

	e.Value.([]RenderComponent)[sliceIndex] = element
}

// Puts the cv.Empty value at the given position
func (cv *RenderComponentChunkedVec) DeleteAt(listIndex, sliceIndex uint) {
	cv.PutAt(cv.Empty, listIndex, sliceIndex)
}

// Returns the value that is on the given position
func (cv *RenderComponentChunkedVec) Get(listIndex, sliceIndex uint) RenderComponent {
	e := cv.List.Front()
	for i := uint(0); i < listIndex; e = e.Next() {
		i++
	}

	return e.Value.([]RenderComponent)[sliceIndex]
}

// Adds chunks (list nodes) to the RenderComponentChunkedVec
func (cv *RenderComponentChunkedVec) Grow(n int) {
	if n < 0 {
		panic("Can't grow RenderComponentChunkedVec with a negative amount")
	}

	for i := 0; i < n; i++ {
		slice := make([]RenderComponent, cv.ChunkSize)
		cv.List.PushBack(slice)
	}
}

// Remove list nodes that has arrays that are with the Empty element only
func (cv *RenderComponentChunkedVec) Shrink() {
	for e := cv.List.Front(); e != nil; e = e.Next() {
		allEmpty := true
		for _, value := range e.Value.([]RenderComponent) {
			if value != cv.Empty {
				allEmpty = false
				break
			}
		}

		if allEmpty {
			cv.List.Remove(e)
		}
	}
}

// Returns the number of non-empty valued elements
func (cv *RenderComponentChunkedVec) Len() int {
	number := 0

	for e := cv.List.Front(); e != nil; e = e.Next() {
		for _, value := range e.Value.([]RenderComponent) {
			if value != cv.Empty {
				number++
			}
		}
	}

	return number
}

// Returns the current capacity of the RenderComponentChunkedVec
// i.e. the number of elements it can currently hold without growing
func (cv *RenderComponentChunkedVec) Cap() int {
	return cv.List.Len() * int(cv.ChunkSize)
}

// Iter returns a channel of type RenderComponent that you can range over.
func (cv *RenderComponentChunkedVec) Iter() <-chan RenderComponent {
	ch := make(chan RenderComponent)

	go func() {
		for e := cv.List.Front(); e != nil; e = e.Next() {
			for _, value := range e.Value.([]RenderComponent) {
				ch <- value
			}
			close(ch)
		}
	}()

	return ch
}

// Checks if the RenderComponentChunkedVec contains the given element
func (cv *RenderComponentChunkedVec) Contains(element RenderComponent) bool {
	for e := cv.List.Front(); e != nil; e = e.Next() {
		for _, value := range e.Value.([]RenderComponent) {
			if value == element {
				return true
			}
		}
	}

	return false
}

// Checks if the RenderComponentChunkedVec contains all of the given element
func (cv *RenderComponentChunkedVec) ContainsAll(searchingFor ...RenderComponent) bool {
	for _, s := range searchingFor {
		if !cv.Contains(s) {
			return false
		}
	}

	return true
}

// Checks if this RenderComponentChunkedVec is equal to another one
// Two RenderComponentChunkedVecs are equal if they have the same number of lists
// with slices that have the same values
func (cv *RenderComponentChunkedVec) Equal(other *RenderComponentChunkedVec) bool {
	// no worries, the complexity of this is O(1)
	if cv.List.Len() != other.List.Len() {
		return false
	}

	e2 := other.List.Front()
	for e1 := cv.List.Front(); e1 != nil; e1 = e1.Next() {
		len1 := len(e1.Value.([]RenderComponent))
		len2 := len(e2.Value.([]RenderComponent))
		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if e1.Value.([]RenderComponent)[i] != e2.Value.([]RenderComponent)[i] {
				return false
			}
		}

		e2 = e2.Next()
	}

	return true
}

// Clone returns a clone of the RenderComponentChunkedVec.
// Does NOT clone the underlying elements.
func (cv *RenderComponentChunkedVec) Clone() *RenderComponentChunkedVec {
	clonedRenderComponentChunkedVec := NewRenderComponentChunkedVec(cv.ChunkSize)

	var listIndex uint = 0
	for e := cv.List.Front(); e != nil; e = e.Next() {
		for index, value := range e.Value.([]RenderComponent) {
			clonedRenderComponentChunkedVec.PutAt(value, listIndex, uint(index))
		}

		listIndex++
	}

	return clonedRenderComponentChunkedVec
}

// Clears all the data in the RenderComponentChunkedVec
func (cv *RenderComponentChunkedVec) Clear() {
	for e := cv.List.Front(); e != nil; e = e.Next() {
		cv.List.Remove(e)
	}
}

func (cv *RenderComponentChunkedVec) String() string {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "{\n")
	for e := cv.List.Front(); e != nil; e = e.Next() {
		slice := e.Value.([]RenderComponent)
		if _, err := fmt.Fprintf(&buff, fmt.Sprintf("\t%s\n", slice)); err != nil {
			panic("Can't write to buffer")
		}
	}
	fmt.Fprintf(&buff, "\n}")

	return buff.String()
}
