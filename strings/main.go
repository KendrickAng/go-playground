package main

import "fmt"

func main() {
	one()
}

func a(s string) {
	fmt.Println("#################################### s ####################################")
}

func one() {
	a("one")
	// strings are just read-only slices of bytes with extra syntactic support
	// slices with no capacity (read-only)
	slash := "/usr/ken"[0]
	fmt.Println(slash)

	// since the array underlying a string is hidden from view, any conversions
	// must first copy the array; we can't access it
	byts := []byte("hello world")
	sbyts := string(byts)
	byts[0] = 'f'
	fmt.Println(string(byts))
	fmt.Println(sbyts)
}