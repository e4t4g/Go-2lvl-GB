package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	const counter int = 1000
	var result int64
	var workers = make(chan struct{}, counter)

	for i := 1; i < counter; i++ {
		workers <- struct{}{}

		go func(){
			atomic.AddInt64(&result, 1)
		}()
	}
	time.Sleep(time.Duration(1)*time.Millisecond)
	fmt.Println(result)
}
