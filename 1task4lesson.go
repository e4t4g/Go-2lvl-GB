package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ch:=make(chan os.Signal,1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)


	go func() {
		for sig := range ch {
			switch sig {
			case syscall.SIGTERM:
				fmt.Println("SIGTERM received")
				os.Exit(0)
			case syscall.SIGINT:
				fmt.Println("SIGINT received")
				os.Exit(0)
			}
		}
	}()

	func() {
		for {
			fmt.Println(time.Now().Format("15:04:05 "))
			time.Sleep(1*time.Second)
		}
	}()
}