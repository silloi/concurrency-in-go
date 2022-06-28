package main

import (
	"fmt"
)

func main() {
	// // batch oriented
	// multiply := func(values []int, multiplier int) []int {
	// 	multipliedValues := make([]int, len(values))
	// 	for i, v := range values {
	// 		multipliedValues[i] = v * multiplier
	// 	}
	// 	return multipliedValues
	// }

	// add := func(values []int, additive int) []int {
	// 	addedValues := make([]int, len(values))
	// 	for i, v := range values {
	// 		addedValues[i] = v + additive
	// 	}
	// 	return addedValues
	// }

	// ints := []int{1, 2, 3, 4}
	// for _, v := range add(multiply(ints, 2), 1) {
	// 	fmt.Println(v)
	// }

	// for _, v := range multiply(add(multiply(ints, 2), 1), 2) {
	// 	fmt.Println(v)
	// }

	// // imperative
	// for _, v := range ints {
	// 	fmt.Println(2*v*2 + 1)
	// }

	// // stream oriented
	// multiply := func(value, multiplier int) int {
	// 	return value * multiplier
	// }

	// add := func(value, additive int) int {
	// 	return value + additive
	// }

	// ints := []int{1, 2, 3, 4}
	// for _, v := range ints {
	// 	fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	// }

	// // rewrite with channel
	// generator := func(done <-chan interface{}, integers ...int) <-chan int {
	// 	intStream := make(chan int, len(integers))
	// 	go func() {
	// 		defer close(intStream)
	// 		for _, i := range integers {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case intStream <- i:
	// 			}
	// 		}
	// 	}()
	// 	return intStream
	// }

	// multiply := func(
	// 	done <-chan interface{},
	// 	intStream <-chan int,
	// 	multiplier int,
	// ) <-chan int {
	// 	multipliedStream := make(chan int)
	// 	go func() {
	// 		defer close(multipliedStream)
	// 		for i := range intStream {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case multipliedStream <- i * multiplier:
	// 			}
	// 		}
	// 	}()
	// 	return multipliedStream
	// }

	// add := func(
	// 	done <-chan interface{},
	// 	intStream <-chan int,
	// 	additive int,
	// ) <-chan int {
	// 	addedStream := make(chan int)
	// 	go func() {
	// 		defer close(addedStream)
	// 		for i := range intStream {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case addedStream <- i + additive:
	// 			}
	// 		}
	// 	}()
	// 	return addedStream
	// }

	// done := make(chan interface{})
	// defer close(done)

	// intStream := generator(done, 1, 2, 3, 4)
	// pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	// for v := range pipeline {
	// 	fmt.Println(v)
	// }

	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	// done := make(chan interface{})
	// defer close(done)

	// for num := range take(done, repeat(done, 1), 10) {
	// 	fmt.Printf("%v ", num)
	// }

	// repeatFn := func(
	// 	done <-chan interface{},
	// 	fn func() interface{},
	// ) <-chan interface{} {
	// 	valueStream := make(chan interface{})
	// 	go func() {
	// 		defer close(valueStream)
	// 		for {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case valueStream <- fn():
	// 			}
	// 		}
	// 	}()
	// 	return valueStream
	// }

	// done := make(chan interface{})
	// defer close(done)
	// rand := func() interface{} { return 1 } // { return rand.Int() }

	// for num := range take(done, repeatFn(done, rand), 10) {
	// 	fmt.Println(num)
	// }

	toString := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	done := make(chan interface{})
	defer close(done)

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}

	fmt.Printf("message: %s...", message)
}
