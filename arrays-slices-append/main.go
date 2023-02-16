package main

import (
	"bytes"
	"fmt"
)

func main() {
	one()
	two()
	three()
	four()
	five()
	six()
}

func a(s string) {
	fmt.Println("################################ " + s + " ################################")
}

func one() {
	a("one")
	// arrays are fixed size contiguous chunks of memory.
	// the size of an array is part of its type - limiting its expressive power
	var buffer [256]byte
	buffer[0] = 'a'
	fmt.Println(len(buffer))

	// a slice is a data structure describing a contiguous section of an array
	// , which is stored separately from the slice variable.
	var slice []byte = buffer[100:150]

	// the slice variable is just a 'slice header' looking like this.
	// a struct value holding a pointer and length, NOT a pointer to a struct.
	type sliceHeader struct {
		Length int
		// Also capacity, but see later
		// Capacity      int
		ZerothElement *byte
	}

	// a slice can be sliced. equivalent to 105 through 109 of original array
	slice2 := slice[5:10]

	// slice header is passed by value.
	// the contents of a slice argument can be modified by the fn, but its header cannot.
	subtractOne := func(s []byte) []byte {
		s = s[0 : len(s)-1]
		return s
	}
	fmt.Println("Before: len(slice) =", len(slice2))
	newSlice := subtractOne(slice2)
	fmt.Println("After: len(slice) =", len(slice2))
	fmt.Println("After:  len(newSlice) =", len(newSlice))
}

func two() {
	a("two")
	// use pointers to modify the slice header.
	_ = func(slicePtr *[]byte) {
		slice := *slicePtr
		*slicePtr = slice[0 : len(slice)-1]
	}

	// idiomatic to use pointer receivers to modify a slice
	type path []byte

	_ = func(p *path) {
		i := bytes.LastIndex(*p, []byte("/"))
		if i >= 0 {
			*p = (*p)[0:i]
		}
	}
	// when modify the slice in-place, value receiver is sufficient.
	_ = func(p path) {
		for i, _ := range p {
			p[i] = 'a'
		}
	}
}

func three() {
	a("three")
	// slice header also has capacity - how much space the underlying array has.
	type sliceHeader struct {
		Length        int
		Capacity      int
		ZerothElement *byte
	}

	// trying to grow a slice beyond its capacity causes panic
	var arr [5]byte
	// arr2 = arr[0:6]
	arr2 := arr[0:2]
	fmt.Println("len", len(arr2), "cap", cap(arr2))

	// make allocaetes a new array and creates a slice header to describe it.
	// so we don't need to use new() then copy the data over
	// var aa [5]byte
	aaa := new([]byte)
	fmt.Printf("%#v\n", aaa)
	fmt.Printf("%+v\n", aaa)
	// (*aaa)[1] = 'a'
}

func four() {
	a("four")
	// golang treats a nil slice as an empty slice
	arr := []int(nil)
	arr = append(arr, 1)
	fmt.Printf("%#v\n", arr)

	type sliceHeader struct {
		Length int
		Capacity int
		ZerothElement *int
	}
	// this is the nil slice
	_ = sliceHeader{
		Length: 0,
		Capacity: 0,
		ZerothElement: nil,
	}
	// same thing
	_ = sliceHeader{}
}

func five() {
	a("five")
	// make(type, length, capacity)
	// length argument defaults to the capacity - length 10, cap 10
	arr := make([]int, 10)
	fmt.Println("len", len(arr), "cap", cap(arr))
}

func six() {
	a("six")
	arr := []int{1, 2, 3}
	// arr[:] basically copies the slice header; underlying array is shared.
	arr2 := arr[:]
	arr2[0] = 3
	fmt.Printf("arr: %#v\n", arr)
	fmt.Printf("arr2: %#v\n", arr2)

	// copy performs a shallow copy of the slice
	arr = []int{1, 2, 3, 4}
	arr2 = make([]int, 4, len(arr))
	copy(arr2, arr)
	arr2[0] = 4
	fmt.Printf("arr: %#v\n", arr)
	fmt.Printf("arr2: %#v\n", arr2)
}

func Extend(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

// Append appends the items to the slice.
// First version: just loop calling Extend.
func Append(slice []int, items ...int) []int {
	for _, item := range items {
		slice = Extend(slice, item)
	}
	return slice
}

func seven() {
	a("seven")
	// this is how append looks like
	slice1 := []int{0, 1, 2, 3, 4}
	slice2 := []int{55, 66, 77}
	fmt.Println(slice1)
	slice1 = Append(slice1, slice2...) // The '...' is essential!
	fmt.Println(slice1)
}
