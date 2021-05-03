package main


import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sort"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

		const n int = 100
		wg := sync.WaitGroup{}
		var (
			lock sync.Mutex
			sl   = make([]int, 0)
		)


		for i := 0; i < n; i++ {
			wg.Add(1)
			go func(i int, lock *sync.Mutex) {
				lock.Lock()
				defer lock.Unlock()
				defer wg.Done()
				if i%3 == 0 {
					sl = append(sl, i)
				} else {
				}
			}(i, &lock)
		}

		runtime.GC()

		wg.Wait()
		sort.Ints(sl)
		fmt.Printf("%d \n", sl)

	}

