package main

func main() {
	// see or-channel

	// for val := range myChan {
	// 	// valに対して何かする
	// }

	// loop:
	// for {
	// 	select {
	// 	case <-done:
	// 		break loop
	// 	case maybeVal, ok := <-myChan:
	// 		if ok == false {
	// 			return // あるいはforからbreakするとか
	// 		}
	// 		// valに対して何かする
	// 	}
	// }

	// orDone := func(done, c <-chan interface{}) <-chan interface{} {
	// 	valStream := make(chan interface{})
	// 	go func() {
	// 		defer close(valStream)
	// 		for {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case v, ok := <-c:
	// 				if ok == false {
	// 					return
	// 				}
	// 				select {
	// 				case valStream <- v:
	// 				case <-done:
	// 				}
	// 			}
	// 		}
	// 	}()
	// 	return valStream
	// }

	// for val := range orDone(done, myChan) {
	// 	// valに対して何かする
	// }
}
