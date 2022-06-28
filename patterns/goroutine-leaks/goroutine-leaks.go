package main

import (
	"fmt"
	"time"
)

func main() {
	// doWork := func(strings <-chan string) <-chan interface{} {
	// 	completed := make(chan interface{})
	// 	go func() {
	// 		defer fmt.Println("doWork exited.")
	// 		defer close(completed)
	// 		for s := range strings {
	// 			// 何かおもしろい処理
	// 			fmt.Println(s)
	// 		}
	// 	}()
	// 	return completed
	// }

	// doWork(nil)
	// // もう少し何かしらの処理がここで行われる
	// fmt.Println("Done.")

	// doWork := func(
	// 	done <-chan interface{},
	// 	strings <-chan string,
	// ) <-chan interface{} {
	// 	terminated := make(chan interface{})
	// 	go func() {
	// 		defer fmt.Println("doWork exited.")
	// 		defer close(terminated)
	// 		for {
	// 			select {
	// 			case s := <-strings:
	// 				// 何かおもしろい処理
	// 				fmt.Println(s)
	// 			case <-done:
	// 				return
	// 			}
	// 		}
	// 	}()
	// 	return terminated
	// }

	// done := make(chan interface{})
	// terminated := doWork(done, nil)

	// go func() {
	// 	// 1秒後に操作をキャンセルする
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("Canceling doWork goroutine...")
	// 	close(done)
	// }()

	// <-terminated
	// fmt.Println("Done.")

	// newRandStream := func() <-chan int {
	// 	randStream := make(chan int)
	// 	go func() {
	// 		defer fmt.Println("newRandStream closure exited.")
	// 		defer close(randStream)
	// 		for {
	// 			randStream <- 12345 // rand.Int()
	// 		}
	// 	}()

	// 	return randStream
	// }

	// randStream := newRandStream()
	// fmt.Println("3 random ints:")
	// for i := 1; i <= 3; i++ {
	// 	fmt.Printf("%d: %d\n", i, <-randStream)
	// }

	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- 12345: // rand.Int()
				case <-done:
					return
				}

			}
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	// 処理が実行中であることをシミュレート
	time.Sleep(1 * time.Second)
}
