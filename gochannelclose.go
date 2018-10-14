package main

func main() {
	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다
	close(ch)

	// 채널 수신
	println(<-ch)
	println(<-ch)

	if _, success := <-ch; !success {
		println("더이상 데이타 없음.")
	}
}
