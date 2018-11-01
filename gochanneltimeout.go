package main

import "time"

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func(done chan bool) {
		time.Sleep(5 * time.Second)
		done <- true
	}(ch1)

	go func(done chan bool) {
		time.Sleep(10 * time.Second)
		done <- true
	}(ch2)

	// time.After()는 입력파라미터에 지정된 시간이
	// 지나면 Ready되는 채널을 리턴한다
	timeoutChan := time.After(1 * time.Second)

	select {
	case <-ch1:
		println("run1")

	case <-ch2:
		println("run2")

		// select 문 내에서 타임아웃 체크
	case <-timeoutChan:
		println("timeout")
	}
}
