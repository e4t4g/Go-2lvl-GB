package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	const n int = 10
	wg := sync.WaitGroup{}
	sl := make([]int, 0)
	var lock sync.Mutex
	var j int

	//Implementation of Mutex
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int, lock *sync.Mutex) {
			defer wg.Done()
			defer lock.Unlock()
			lock.Lock()
			sl = append(sl, i)
		}(i, &lock)
	}
	wg.Wait()
	sort.Ints(sl)
	fmt.Printf("%d \n", sl)

	//Implementation of Wait Group
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			j += 1
			fmt.Printf("%d \n", j)
		}(i)
	}
	wg.Wait()

	fmt.Print("main goroutine is done")
}