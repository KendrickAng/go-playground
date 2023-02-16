package main

import (
	"fmt"
	"time"
)

// https://go.dev/blog/defer-panic-and-recover
func main() {
	one()
	two()
	fmt.Println(three())
	// four()
	five()
}

func announce(s string) {
	fmt.Println("########################## " + s + " ##########################")
}

func one() {
	announce("one")
	// 1. A deferred function's arguments are evaluated when the defer statement is evaluated.
	i := 0
	defer fmt.Println(i) // 0
	i++
}

func two() {
	announce("two")
	// Deferred function calls are executed in LIFO order after the surrounding fn returns
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func three() (i int) {
	announce("three")
	// Deferred functions may read/write to the returning function's named return values
	// Convenient for modifying error reutrn value of a function
	defer func() { i++ }()
	return 1
}

func four() {
	announce("four")
	// when a function F calls panic, execution of F stops, any deferred functions in F are
	// executed normally, then F returns to its caller.
	// In the caller, F then behaves like a call to panic, then the process continues up the
	// stack until all functions in the current goroutine returns, then the program crashes.
	go func() { panic("help") }()
	defer fmt.Println("deferring")
	time.Sleep(3 * time.Second)
}

func five() {
	announce("five")
	// recover regains control of a panicking goroutine. It only works in deferred functions. 
	// In normal execution, the call to recover returns nil.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f:", r)
		}
	}()
	panic("panicking!")
}