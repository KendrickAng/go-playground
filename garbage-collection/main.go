package main

import (
	"fmt"
	"runtime"
)

// This one doesn't work as intended.
// The GC should wait forever for the long-running process before starting (stop-world GC)
// But it doesn't! 
func main() {
	go func() {
		for i := 0; i <= 1000_000_000_000; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("Dropping mic")
	// yield execution to force execute other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
}
