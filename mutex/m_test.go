package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

var (
	globalMap   = map[int]int{}
	globalMapMu = sync.RWMutex{}
)

func main() {
	CheckMutex(0.3)
	CheckRWMutex(0.3)
	fmt.Println("main done")
}

func CheckMutex(n float32) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			if rand.Float32() < n {
				globalMapMu.Lock()
				globalMap[i] = i
				globalMapMu.Unlock()
				return
			}
			globalMapMu.Lock()
			//fmt.Print(globalMap[i])
			globalMapMu.Unlock()
		}(i)
	}
	wg.Wait()
}

func CheckRWMutex(n float32) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if rand.Float32() < n {
				globalMapMu.Lock()
				globalMap[i] = i
				globalMapMu.Unlock()
				return
			}
			globalMapMu.RLock()
			//fmt.Print(globalMap[i])
			globalMapMu.RUnlock()
		}(i)
	}
	wg.Wait()
}

func BenchmarkMutex10_90(b *testing.B) {
	for i:=0; i <b.N; i++ {
		CheckMutex(0.9)
	}
}

func BenchmarkMutex50_50(b *testing.B) {
	for i:=0; i <b.N; i++ {
		CheckMutex(0.5)
	}
}

func BenchmarkMutex90_10(b *testing.B) {
	for i:=0; i <b.N; i++ {
		CheckMutex(0.1)
	}
}

func BenchmarkRWMutex10_90(b *testing.B) {
	for i:=0; i <b.N; i++ {
		CheckRWMutex(0.9)
	}
}

func BenchmarkRWMutex50_50(b *testing.B) {
	for i:=0; i <b.N; i++ {
		CheckRWMutex(0.5)
	}
}

func BenchmarkRWMutex90_10(b *testing.B) {
	for i:=0; i <b.N; i++ {
		CheckRWMutex(0.1)
	}
}