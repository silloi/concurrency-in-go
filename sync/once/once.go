package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once

	// var increments sync.WaitGroup
	// increments.Add(100)
	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		defer increments.Done()
	// 		once.Do(increment)
	// 	}()
	// }

	// increments.Wait()
	// fmt.Printf("Count is %d\n", count) // 1

	count = 0 // var count int
	// increment := func() { count++ }
	decrement := func() { count-- }

	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count) // 1

	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }
	onceA.Do(initA) // dead lock!
}
