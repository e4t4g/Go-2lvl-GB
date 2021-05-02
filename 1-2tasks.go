package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	const n int = 10
	wg := sync.WaitGroup{}
	var j int
	var (
		lock sync.Mutex
		sl = make([]int, 0)
	)
	
	//Implementation of Mutex
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int, lock *sync.Mutex) {
			lock.Lock()
			defer lock.Unlock()
			defer wg.Done()
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

