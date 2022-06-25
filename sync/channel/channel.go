package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

func main() {
	// var receiveChan <-chan interface{}
	// var sendChan chan<- interface{}
	// dataStream := make(chan interface{})

	// 正しい記述
	// receiveChan = dataStream
	// sendChan = dataStream

	// intStream := make(chan int)

	stringStream := make(chan string)
	go func() {
		// 読み込み
		stringStream <- "Hello channels!"
	}()
	fmt.Println(<-stringStream)

	// writeStream := make(chan<- interface{})
	// readStream := make(<-chan interface{})

	// <-writeStream()
	// readStream <- struct{}{}

	// go func() {
	// 	if 0 != 1 { // deadlock
	// 		return
	// 	}
	// 	stringStream <- "Hello channels!"
	// }()
	// fmt.Println(<-stringStream)

	go func() {
		stringStream <- "Hello channels!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salutation)

	valueStream := make(chan interface{})
	close(valueStream)

	// intStream := make(chan int)
	// close(intStream)
	// integer, ok := <-intStream
	// fmt.Printf("(%v): %v", ok, integer)

	// intStream := make(chan int)
	// go func() {
	// 	defer close(intStream)
	// 	for i := 1; i <= 5; i++ {
	// 		intStream <- i
	// 	}
	// }()

	// for integer := range intStream {
	// 	fmt.Printf("%v ", integer)
	// }

	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()

	// var dataStream chan interface{}
	// dataStream = make(chan interface{}, 4)

	// a := make(chan int)
	// b := make(chan int, 0)

	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
