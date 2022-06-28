package main

func main() {
	for { // 無限ループまたは何かのイテレーションを回す
		select {
		// チャネルに対して何かを行う
		}
	}

	done := make(chan interface{})
	stringStream := make(chan string)

	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream <- s:
		}
	}

	for {
		select {
		case <-done:
			return
		default:
		}

		// 割り込みできない処理をする
	}

	for {
		select {
		case <-done:
			return
		default:
			// 割り込みできない処理をする
		}
	}
}
