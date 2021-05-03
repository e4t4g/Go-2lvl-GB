package main

import (
	"sync"
)

func main() {
	arr := map[string]int{}
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i <10; i++ {
			arr["a"] = i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i <10; i++ {
			arr["g"] = i + 1
		}
	}()

	wg.Wait()
}