package main

import "fmt"

func main() {
	prefix := "Hello"
	// go function closures are evaluated at RUNTIME, not compile time.
	f := func(s string) {
		fmt.Println(prefix + " World!")
	}
	// Hello World
	f(prefix)
	prefix = "Goodbye"
	// Expected: Hello World
	// Got: Goodbye World
	f(prefix)
}
