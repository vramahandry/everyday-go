package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		printLater("Hello\n", time.Millisecond*100)
	}()
	go func() {
		defer wg.Done()
		printLater("World\n", time.Millisecond*100)
	}()
	go func() {
		defer wg.Done()
		printLater(os.Getenv("USER")+"\n", time.Millisecond*100)
	}()
	wg.Wait()
}
func printLater(msg string, duration time.Duration) {
	time.Sleep(duration)

	fmt.Printf(msg)
}
