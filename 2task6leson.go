package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()
		for i:=0; i<10; i++ {
			fmt.Printf("%d",i)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c:='a'; c <= 'z'; c+=1 {
			fmt.Printf("%c",c)
		}
	}()

	runtime.Gosched()
	wg.Wait()
}
