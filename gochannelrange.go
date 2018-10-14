package main

func main() {
	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다
	close(ch)

	// 방법1
	// 채널이 Close된 때 까지 계속 수신
	/*
		for {
			if i, success := <-ch; success {
				println(i)
			} else {
				break
			}
		}
	*/

	// 방법2
	// 위 표현과 동일한 채널 range 문
	for i := range ch {
		println(i)
	}
}
