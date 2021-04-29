/*package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for sig := range ch {
			switch sig {
			case syscall.SIGTERM, syscall.SIGINT:
				fmt.Println("SIGTERM received")
				<-time.After(1 * time.Second)
				os.Exit(0)
			}
		}
	}()

	func() {
		for {
			fmt.Println(time.Now().Format("15:04:05 "))
			time.Sleep(1 * time.Second)
		}
	}()
}
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch:=make(chan os.Signal,1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()


	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				for {
					fmt.Println(time.Now().Format("15:04:05 "))
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()

	sign := <- ch
	fmt.Println("SIGTERM received", sign)
}


